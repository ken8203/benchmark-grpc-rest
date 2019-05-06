package main

import (
	"log"

	"github.com/ken8203/benchmark-grpc-rest/pkg/grpc"
)

func main() {
	log.Println("starting gRPC server...")
	grpc.RunServer()
}
