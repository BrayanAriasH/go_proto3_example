package main

import (
	"context"
	"log"
	"time"

	pb "github.com/BrayanAriasH/go_proto3_example/cmd/app/hello/hello_grpc"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// if len(os.Args) != 3 {
	// 	log.Fatalf("2 params are required")
	// }
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("error connecting grpc client: %v", err)
	}
	defer conn.Close()
	grpcClient := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := grpcClient.SayHello(ctx, &pb.HelloRequest{Name: "Juan Carlos", Age: 31})
	if err != nil {
		log.Fatalf("error calling grpc server: %v", err)
	}
	log.Printf("grpc response %s", response.Greeting)

}
