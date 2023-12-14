package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	//The name of your project in go.mod
	pb "simpleAPI/gen/proto"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func main() {

	listner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listner)
	if err != nil {
		log.Println(err)

	}
}
