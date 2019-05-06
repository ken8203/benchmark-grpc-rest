package grpc

import (
	"context"
	"testing"

	api "github.com/ken8203/benchmark-grpc-rest/cmd/grpc/proto"
	"google.golang.org/grpc"
)

func init() {
	go RunServer()
}

func BenchmarkGRPC(b *testing.B) {
	cc, err := grpc.Dial("0.0.0.0:8000", grpc.WithInsecure())
	if err != nil {
		b.Fatalf("grpc connection failed %v", err)
	}

	client := api.NewGreeterServiceClient(cc)

	for i := 0; i < b.N; i++ {
		req := &api.HelloRequest{
			Name: "Jay Chung",
		}

		_, err := client.SayHello(context.Background(), req)
		if err != nil {
			b.Fatalf("grpc request failed: %v", err)
		}
	}
}
