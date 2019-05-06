package main

import (
	"log"

	httprest "github.com/ken8203/benchmark-grpc-rest/pkg/http"
)

func main() {
	log.Println("starting RESTful server...")
	httprest.RunServer()
}
