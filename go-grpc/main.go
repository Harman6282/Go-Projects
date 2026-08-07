package main

import (
	"grpc-crud-go/db"
	"grpc-crud-go/server"
	"log"
	"net"

	pb "grpc-crud-go/grpc-crud-go/proto"

	"google.golang.org/grpc"
)


func main () {
	err := db.InitMongo()
	if err != nil {
		log.Fatal("Mongodb connection failed", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server.UserServer{})
	log.Println("grpc server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to start grpc server")
	}
}