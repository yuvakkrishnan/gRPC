package main

import (
	"log"
	"net"

	pb "github.com/yuvakkrishnan/grpc-go/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Sever struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Failed to listen on:%v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to server:%v\n", err)
	}
}
