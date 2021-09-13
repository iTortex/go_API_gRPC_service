package main

import (
	"context"
	// "errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/url"
	"os"

	pb "example.com/1module/modulegrpc" // pb = protobuf  в начале мы настроили среду в module после чего указываем откуда мы подтягиваем функции (возможно неверное описание)
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

type user struct {
	ID       int32
	URL      string
	ShortUrl string
}
type UserManagmentServer struct {
	conn *pgx.Conn
	pb.UnimplementedUserManagmentServer
}

func Shorting() (string, error) {
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
		ID SERIAL PRIMARY KEY,
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

	rows := s.conn.QueryRow(context.Background(), "select * from urls where URL=$1", in.GetName())
	var some_user user

	err = rows.Scan(&some_user.ID, &some_user.URL, &some_user.ShortUrl)
	if err != nil && err.Error() != "no rows in result set" {
			log.Printf("error check: %v\n", err)
			return nil, err
		}
	if in.GetName() == some_user.URL {
			log.Printf("Returned: %v", some_user.ShortUrl)
			return &pb.ShortURL{Shortname: some_user.ShortUrl}, nil
	}

	str, err := Shorting()
	if err != nil {
		log.Printf("Cant short it: %v", err)
		return nil, err
	}
	rows = s.conn.QueryRow(context.Background(), "select * from urls where ShortURL=$1", str)
	err = rows.Scan(&some_user.ID, &some_user.URL, &some_user.ShortUrl)
	if err != nil && err.Error() != "no rows in result set" {
		log.Printf("error check: %v\n", err)
		return nil, err
	}
	for err == nil {
		str, err = Shorting()
		if err != nil {
			log.Printf("Cant short it: %v", err)
			return nil, err
		}
		rows = s.conn.QueryRow(context.Background(), "select * from urls where ShortURL=$1", str)
		err = rows.Scan(&some_user.ID, &some_user.URL, &some_user.ShortUrl)
		if err != nil && err.Error() != "no rows in result set" {
			log.Printf("error check: %v\n", err)
			return nil, err
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
	log.Printf("Returned END: %v", str)
	return &pb.ShortURL{Shortname: str}, nil
}

func (s *UserManagmentServer) Get(ctx context.Context, in *pb.ShortURL) (*pb.URL, error) {
	
	log.Printf("Received: %v", in.GetShortname())
	rows := s.conn.QueryRow(context.Background(), "select * from urls where ShortURL=$1", in.GetShortname())
	log.Printf("rows: %v\n", rows) 
	var some_user user

	if err := rows.Scan(&some_user.ID, &some_user.URL, &some_user.ShortUrl); err != nil {
		return nil, err
	}
	return &pb.URL{Name: some_user.URL}, nil
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
	database_url := "postgres://amarcele:qwertyui@db:5432/samplegres"
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
