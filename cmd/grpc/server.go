package grpc

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	api "github.com/ken8203/benchmark-grpc-rest/cmd/grpc/proto"
	"google.golang.org/grpc"
)

type GreeterServer struct{}

func RunServer() {
	lis, _ := net.Listen("tcp", ":8000")

	srv := grpc.NewServer()
	api.RegisterGreeterServiceServer(srv, &GreeterServer{})
	log.Fatalln(srv.Serve(lis))
}

func (g *GreeterServer) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloResponse, error) {
	return &api.HelloResponse{
		Message: fmt.Sprintf("Hey there %s %f", req.GetName(), rand.Float64()),
	}, nil
}
