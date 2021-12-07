package server

import (
	"context"
	"log"
	"net"

	sumpb "github.com/mjsjinsu/proto-go-sum/protos/sum"
	"google.golang.org/grpc"
)

const portNumber = "9000"

type sumServer struct {
	sumpb.GRPCFuncsListServer
}

func (s *sumServer) Sum(ctx context.Context, req *sumpb.SumArgs) (*sumpb.SumReturns, error) {
	log.Println(req)
	calc := req.Value1 + req.Value2
	return &sumpb.SumReturns{Value: calc}, nil
}

func main() {
	ltn, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	sumpb.RegisterGRPCFuncsListServer(grpcServer, &sumServer{})
	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(ltn); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
