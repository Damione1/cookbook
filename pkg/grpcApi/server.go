package grpcApi

import (
	"fmt"

	"github.com/Damione1/portfolio-playground/pkg/pb"
	"github.com/Damione1/portfolio-playground/pkg/token"
	"github.com/Damione1/portfolio-playground/util"
)

type Server struct {
	pb.UnimplementedPortfolioServiceServer
	config     util.Config
	tokenMaker token.Maker
}

func NewServer(config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create token maker. %v", err)
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
