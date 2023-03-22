package grpcApi

import (
	"context"

	"github.com/Damione1/portfolio-playground/db/models"
	"github.com/Damione1/portfolio-playground/pkg/pb"
	"github.com/Damione1/portfolio-playground/pkg/pbx"
	"github.com/Damione1/portfolio-playground/util"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	err := validateCreateUserRequest(req)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	user := models.User{
		Email:    req.GetEmail(),
		Password: hashedPassword,
		Name:     req.GetName(),
	}

	ok, err := user.Exists(ctx, boil.GetContextDB())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check if user exists: %s", err)
	}
	if ok {
		return nil, status.Errorf(codes.AlreadyExists, "user with email %s already exists", req.GetEmail())
	}

	err = user.Insert(ctx, server.config.DB, boil.Infer())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert user: %s", err)
	}

	pbUser := pbx.DbUserToProto(&user)

	return &pb.CreateUserResponse{
		User: pbUser,
	}, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) error {
	//validate request
	return nil
}
