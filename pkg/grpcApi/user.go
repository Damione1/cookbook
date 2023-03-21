package grpcApi

import (
	"context"

	"github.com/Damione1/portfolio-playground/pkg/pb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	//email := req.GetEmail()
	//hashedPassword, err := util.HashPassword(req.GetPassword())
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal, "failed to hash password. %v", err)
	//}
	//name := req.GetName()
	//
	//user, err := models.Users(models.BlogPostTableColumns.Email.EQ(email)).One(ctx, server.db)
	//if err != nil {
	//	return nil, err
	//}
	return &pb.CreateUserResponse{}, nil
}
