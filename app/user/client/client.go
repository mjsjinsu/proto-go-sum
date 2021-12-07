package main

import (
	"context"
	"fmt"
	"log"
	"time"

	userpb "github.com/mjsjinsu/proto-go-sum/protos/v1/user"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:9000"
)

func main() {
	fmt.Println("Hello Client")
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect : %v", err)
	}

	defer conn.Close()

	c := userpb.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetUser(ctx, &userpb.GetUserRequest{UserId: "1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.UserMessage)
}
