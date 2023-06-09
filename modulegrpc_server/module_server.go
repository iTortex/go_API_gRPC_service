package main

import (
	"context"
	"log"
	"math/rand"
	"errors"
	"net"
	"net/url"
	"os"
	"bufio"
	"time"
	"fmt"

	pb "example.com/1module/modulegrpc"
	"google.golang.org/grpc"
	"github.com/jackc/pgx/v4"
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

var in_memory = make(map[string]string)
var check_shorts = make(map[string]string)

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
	if os.Getenv("DB_USAGE") == "true" { 
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
		log.Printf("Received: %v\n", in.GetName())
		if err := Validate(in.GetName()); err != nil {
			log.Printf("Not valid URL: %v\n", err)
			return nil, err
		}

		rows := s.conn.QueryRow(context.Background(), "select * from urls where URL=$1", in.GetName())
		var some_user user

		err = rows.Scan(&some_user.ID, &some_user.URL, &some_user.ShortUrl)
		if err != nil && err.Error() != "no rows in result set" {
				log.Printf("Error check: %v\n", err)
				return nil, err
			}
		if in.GetName() == some_user.URL {
				log.Printf("Returned: %v\n", some_user.ShortUrl)
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
			log.Printf("Error check: %v\n", err)
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
				log.Printf("Error check: %v\n", err)
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
		log.Printf("Returned: %v\n", str)
		return &pb.ShortURL{Shortname: str}, nil
	} else {

		name := in.GetName()
		log.Printf("Received: %v\n", name)
		if err := Validate(name); err != nil {
			log.Printf("Not valid URL: %v\n", err)
			return nil, err
		}

		if _, ok := in_memory[name]; ok {
			return &pb.ShortURL{Shortname: in_memory[name]}, nil
		}
		str, err := Shorting()
		if err != nil {
			log.Printf("Cant short it: %v", err)
			return nil, err
		}

		for _, ok := check_shorts[str]; ok; {
		str, err = Shorting()
			if err != nil {
				log.Printf("Cant short it: %v", err)
				return nil, err
			}
		}

		check_shorts[str] = name
		in_memory[name] = str

		log.Printf("Returned: %v\n", str)

		file, err := os.OpenFile("names.txt", os.O_APPEND | os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("Cant find old data: %v", err)
			return nil, err
		}
		defer file.Close()

		new_member := name + " " + str + "\n"
		file.WriteString(new_member)

		return &pb.ShortURL{Shortname: in_memory[name]}, nil
	}
}


func (s *UserManagmentServer) Get(ctx context.Context, in *pb.ShortURL) (*pb.URL, error) {
	if os.Getenv("DB_USAGE") == "true" { 
		log.Printf("Received: %v\n", in.GetShortname())
		rows := s.conn.QueryRow(context.Background(), "select * from urls where ShortURL=$1", in.GetShortname())
		var some_user user

		if err := rows.Scan(&some_user.ID, &some_user.URL, &some_user.ShortUrl); err != nil {
			return nil, err
		}
		log.Printf("Returned: %v\n", some_user.URL)
		return &pb.URL{Name: some_user.URL}, nil
	} else {
		log.Printf("Received: %v\n", in.GetShortname())
		tmp := in.GetShortname()
		if _, ok := check_shorts[tmp]; !ok { return nil, errors.New("pridumat' nazvanie")}


		log.Printf("Returned: %v\n", check_shorts[tmp])
		return &pb.URL{Name: check_shorts[tmp]}, nil
		}
}

func (server *UserManagmentServer) Run() error {
	if os.Getenv("DB_USAGE") == "true" {
			lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("Fatal to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterUserManagmentServer(s, server)
		log.Printf("Server listening at %v", lis.Addr())
		rand.Seed(time.Now().UnixNano())
		return s.Serve(lis)
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Fatal to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagmentServer(s, server)
	log.Printf("Server listening at %v", lis.Addr())

	file, err := os.OpenFile("names.txt", os.O_CREATE|os.O_RDONLY, 0777)
	if err != nil {
		log.Printf("Cant find old data: %v", err)
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	for i := 0; i < len(words); i += 2 {
		in_memory[words[i]] = words[i + 1]
		check_shorts[words[i + 1]] = words[i]
	}
	file.Close()
	rand.Seed(time.Now().UnixNano())
	return s.Serve(lis)
}


func main() {
	if os.Getenv("DB_USAGE") == "true" { 
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
	} else {
	var user_mgmt_server *UserManagmentServer = NewUserManagmentServer()
	if err := user_mgmt_server.Run(); err != nil {
		log.Fatalf("Failed to server: %v", err)
		}
	}
}
