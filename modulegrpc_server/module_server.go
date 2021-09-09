package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"net"
	"net/url"

	pb "example.com/module/modulegrpc" // pb = protobuf  в начале мы настроили среду в module после чего указываем откуда мы подтягиваем функции (возможно неверное описание)
	"google.golang.org/grpc"
)

const (
	port  = ":8080"
	bytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
)

var save_data = make(map[string]string)

func NewUserManagmentServer() *UserManagmentServer {
	return &UserManagmentServer{
		//user_list: &pb.UserList{},
	}
}

type UserManagmentServer struct {
	pb.UnimplementedUserManagmentServer
	// 	user_list *pb.UserList // содержит массив указателей юзеров
}

func shorting(URL string) (string, error) {
	b := make([]byte, 10)
	for i := range b {
		b[i] = bytes[rand.Intn(len(bytes))]
	}
	return string(b), nil
}

func validate(URL string) error {
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
	log.Printf("Received: %v", in.GetName())
	if err := validate(in.GetName()); err != nil {
		log.Printf("Not valid URL: %v", err)
		return nil, err
	}
	for short, name := range save_data {
		if name == in.GetName() {
			return &pb.ShortURL{Name: in.GetName(), Shortname: short}, nil
		}
	}
	str, err := shorting(in.GetName())
	if err != nil {
		log.Printf("Cant short it: %v", err)
		return nil, err
	}
	for _, ok := save_data[str]; ok; {
		str, err = shorting(in.GetName())
		if err != nil {
			log.Printf("Cant short it: %v", err)
			return nil, err
		}
	}
	save_data[str] = in.GetName()
	// var short_url int32 = int32(rand.Intn(100))
	// if _, ok := save_data_url[in.GetName()]; ok {
	// 	return save_data_url[in.GetName()], nil
	// }
	// save_data[short_url] = in
	// short := &pb.ShortURL{Name: short_url}
	// save_data_url[in.GetName()] = short
	//created_short_url := &pb.ShortURL{Name: short_url}
	// s.user_list.Users = append(s.user_list.Users, created_user)

	return &pb.ShortURL{Name: in.GetName(), Shortname: str}, nil
}
func (s *UserManagmentServer) Get(ctx context.Context, in *pb.ShortURL) (*pb.URL, error) {
	if _, ok := save_data[in.GetShortname()]; ok {
		return &pb.URL{Name: save_data[in.GetShortname()]}, nil
	}
	err := errors.New("It has not short implementation")
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
	var user_mgmt_server *UserManagmentServer = NewUserManagmentServer()
	if err := user_mgmt_server.Run(); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}