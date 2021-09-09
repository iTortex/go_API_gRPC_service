package main

import (
	"context"
	// "fmt"
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	pb "example.com/module/modulegrpc" // pb = protobuf  в начале мы настроили среду в module после чего указываем откуда мы подтягиваем функции (возможно неверное описание)
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagmentClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 100)
	defer cancel()
	scan := bufio.NewScanner(os.Stdin)
	for {
		log.Print("Please enter the data: ")
		scan.Scan()
		com := strings.Split(string(scan.Text()), " ")
		log.Print(com[0], "  ", com[1])
		if com[0] == "Create" {
			if sh, err := c.Create(ctx, &pb.URL{Name: com[1]}); err == nil {
				log.Printf("Long and short URL is %v",sh.GetShortname())
			} else { 
				log.Printf("Please, try it again")
				continue 
			}
		}
		if com[0] == "Get" {
			if sh, err := c.Get(ctx, &pb.ShortURL{Shortname: com[1]}); err == nil {
				log.Printf("Long URL is %v", sh.GetName())
			} else {
				log.Printf("Please, try it again")
				continue
			}
		}
	}
}