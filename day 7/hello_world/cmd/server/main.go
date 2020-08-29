package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"context"
	"github.com/103cuong/grpc_hello_world/api"
)

type GreeterServiceServer struct {

}

func (s *GreeterServiceServer) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloResponse, error) {
	return &api.HelloResponse{Message: "👋 Hello " + req.Name}, nil
}

func main() {
	const port = ":50000"
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("gRPC server listen failed: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterGreeterServer(grpcServer, &GreeterServiceServer{})
	log.Printf("💅 server ready at 0.0.0.0%s", port)
	grpcServer.Serve(listen)
}
