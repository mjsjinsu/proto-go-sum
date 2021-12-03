package main

import (
	"context"
	"log"
	"net"

	"github.com/mjsjinsu/proto-go-sum/data"
	userpb "github.com/mjsjinsu/proto-go-sum/protos/v1/user"
	"google.golang.org/grpc"
)

const portNumber = "9000"

type userServer struct {
	userpb.UserServer
}

func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userId := req.UserId

	var userMessage *userpb.UserMessage
	for _, u := range data.UserData {
		if u.UserId != userId {
			continue
		}
		userMessage = u
		break
	}
	return &userpb.GetUserResponse{
		UserMessage: userMessage,
	}, nil
}

func (s *userServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {

	userMessages := make([]*userpb.UserMessage, len(data.UserData))
	for i, u := range data.UserData {
		userMessages[i] = u
	}

	return &userpb.ListUsersResponse{
		UserMessages: userMessages,
	}, nil
}

func main() {

	ltn, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServer(grpcServer, &userServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(ltn); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
