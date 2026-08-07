package server

import (
	"context"
	"grpc-crud-go/db"
	pb "grpc-crud-go/grpc-crud-go/proto"
	"grpc-crud-go/models"

	"go.mongodb.org/mongo-driver/bson"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user := models.User{
		UserId:   req.UserId,
		Name:     req.Name,
		Age:      req.Age,
		Email:    req.Email,
		Password: req.Password,
	}

	_, err := db.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		UserId: req.UserId,
		Name:   req.Name,
		Age:    req.Age,
		Email:  req.Email,
	}, nil
}


func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user := models.User{}
	err := db.UserCollection.FindOne(ctx, bson.M{"user_id": req.UserId}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		UserId: req.UserId,
		Name:   req.Name,
		Age:    req.Age,
		Email:  req.Email,
	}, nil 
}



















