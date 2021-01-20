package controllers

import (
	"context"
	"log"

	pb "github.com/dindasigma/go-microservices-user/packages/api/proto/users"
)

type GrpcServer struct{}

func (s GrpcServer) Create(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Print("test", in.FirstName)
	//var user *pb.User
	return &pb.User{FirstName: in.FirstName}, nil
}

func (s GrpcServer) Retrieve(ctx context.Context, in *pb.UserID) (*pb.User, error) {
	log.Print("test", in.Id)
	return &pb.User{Id: in.Id}, nil
}
