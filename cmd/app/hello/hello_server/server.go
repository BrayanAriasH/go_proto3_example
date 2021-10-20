package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/BrayanAriasH/go_proto3_example/cmd/app/hello/hello_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	response := &pb.HelloResponse{}
	response.Greeting = fmt.Sprintf("Hello %s, now we know you're %d years old. Have a great day.", in.Name, in.Age)
	return response, nil
}
func (s *server) HelloEveryOne(*pb.HelloRequest, pb.Greeter_HelloEveryOneServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloEveryOne not implemented")
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error starting server in port %s. error: %v", port, err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})
	log.Printf("Server listening at %v\n", listener.Addr())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Printf("Error starting server %v", err)
	}
}
