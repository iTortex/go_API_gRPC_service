package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"errors"

	pb "example.com/module/modulegrpc" // pb = protobuf  в начале мы настроили среду в module после чего указываем откуда мы подтягиваем функции (возможно неверное описание)
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

var save_data = make(map[int32]*pb.URL)
var save_data_url = make(map[string]*pb.ShortURL)

func NewUserManagmentServer() *UserManagmentServer {
	return &UserManagmentServer{
		//user_list: &pb.UserList{},
	}
}

type UserManagmentServer struct {
	pb.UnimplementedUserManagmentServer
// 	user_list *pb.UserList // содержит массив указателей юзеров
}

func (s *UserManagmentServer) Create(ctx context.Context, in *pb.URL) (*pb.ShortURL, error) {
	log.Printf("Received: %v", in.GetName())
	var short_url int32 = int32(rand.Intn(100))
	if _, ok := save_data_url[in.GetName()]; ok {
		return save_data_url[in.GetName()], nil
	}
	save_data[short_url] = in
	short := &pb.ShortURL{Name: short_url}
	save_data_url[in.GetName()] = short
	//created_short_url := &pb.ShortURL{Name: short_url}
	// s.user_list.Users = append(s.user_list.Users, created_user)
	
	return short, nil
}
func (s *UserManagmentServer) Get(ctx context.Context, in *pb.ShortURL) (*pb.URL, error){
	if _, ok := save_data[in.GetName()]; ok {
		return save_data[in.GetName()], nil
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