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
		if com[0] == "Create" {
			if sh, err := c.Create(ctx, &pb.URL{Name: com[1]}); err == nil {
				log.Printf("Long and short URL is %v : %v", sh.GetName(), sh.GetShortname())
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
		// if err != nil {
		// 	sh, err := c.Create(ctx, &pb.URL{Name: args})
		// 	if err != nil {
		// 		log.Fatalf("could not create short: %v", err)
		// 	}
		// 	log.Printf("Short URL is %v", sh.GetName())
		// } else
		// {
		// 	r, err := c.Get(ctx, &pb.ShortURL{Name: int32(some)})
		// 	if err != nil {
		// 		log.Printf("%v", err)
		// 	} else { log.Printf("Long URL is %v", r.GetName()) }
		// }
	}

// 	for name, age := range new_users {
// 		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
// 		if err != nil {
// 			log.Fatalf("could not create user: %v", err)
// 		}
// 		log.Printf(`User Details:
// NAME: %s
// AGE: %d
// ID: %d`, r.GetName(), r.GetAge(), r.GetId())
// 	}
// 	params := &pb.GetUsersParams{}
// 	r, err := c.GetUsers(ctx, params)
// 	if err != nil {
// 		log.Fatalf("Could not retrive users: %v", err)
// 	}
// 	log.Print("\nUser List : \n")
// 	fmt.Printf("r. GetUsers(): %v\n", r.GetUsers())
}