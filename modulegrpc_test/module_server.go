package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/url"
	"os"

	pb "example.com/module/modulegrpc" // pb = protobuf  в начале мы настроили среду в module после чего указываем откуда мы подтягиваем функции (возможно неверное описание)
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
)

const (
	port  = ":8080"
	bytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
)

func NewUserManagmentServer() *UserManagmentServer {
	return &UserManagmentServer{}
}

type UserManagmentServer struct {
	conn *pgx.Conn
	pb.UnimplementedUserManagmentServer
}

func Shorting(URL string) (string, error) {
	b := make([]byte, 10)
	for i := range b {
		b[i] = bytes[rand.Intn(len(bytes))]
	}
	return string(b), nil
}

func Validate(URL string) error {
	if _, err := url.ParseRequestURI(URL); err != nil {
		return err
	}
	u, err := url.Parse(URL)
	if err != nil || u.Host == "" {
		return err
	}
	return nil
}

func (s *UserManagmentServer) Create(ctx context.Context, in *pb.URL) (*pb.ShortURL, error) {
	createSql := `
	create table if not exists urls(
		URL text,
		ShortURL text
		);
		`
	_, err := s.conn.Exec(context.Background(), createSql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Table creation failed: %v\n", err)
		os.Exit(1)
	}
	log.Printf("Received: %v", in.GetName())
	if err := Validate(in.GetName()); err != nil {
		log.Printf("Not valid URL: %v", err)
		return nil, err
	}

	rows, err := s.conn.Query(context.Background(), "select * from urls")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := [2]string{}
		err := rows.Scan(&user[0], &user[1])
		if err != nil {
			return nil, err
		}
		if in.GetName() == user[0] {
			return &pb.ShortURL{Shortname: user[1]}, nil
		}
	}

	str, err := Shorting(in.GetName())
	for i := 0; i < len(rows.RawValues()); i++ {
		if err != nil {
			log.Printf("Cant short it: %v", err)
			return nil, err
		}
		user := [2]string{}
		err := rows.Scan(&user[0], &user[1])
		if err != nil {
			return nil, err
		}
		if str == user[1] {
			str, err = Shorting(in.GetName())
			if err != nil {
				log.Printf("Cant short it: %v", err)
				return nil, err
			}
		}
	}
	created_short := &pb.ShortURL{Shortname: str}
	tx, err := s.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Beegin failed: %v", err)
	}
	_, err = tx.Exec(context.Background(), "insert into urls(URL, ShortURL) values ($1, $2)", in.GetName(), created_short.Shortname)
	if err != nil {
		log.Fatalf("tx.Exec failed: %v", err)
	}
	tx.Commit(context.Background())
	return &pb.ShortURL{Shortname: str}, nil
}
func (s *UserManagmentServer) Get(ctx context.Context, in *pb.ShortURL) (*pb.URL, error) {
	rows, err := s.conn.Query(context.Background(), "select * from urls")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := [2]string{}
		err := rows.Scan(&user[0], &user[1])
		if err != nil {
			return nil, err
		}
		if in.GetShortname() == user[1] {
			return &pb.URL{Name: user[0]}, nil
		}
	}
	err = errors.New("It has not short implementation")
	return &pb.URL{}, err
}

func (server *UserManagmentServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Fatal to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagmentServer(s, server)
	log.Printf("Server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func main() {
	database_url := "postgres://postgres:mysecretpassword@192.168.99.100:5432/postgres"
	conn, err := pgx.Connect(context.Background(), database_url)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v", err)
	}
	defer conn.Close(context.Background())
	var user_mgmt_server *UserManagmentServer = NewUserManagmentServer()
	user_mgmt_server.conn = conn
	if err := user_mgmt_server.Run(); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
