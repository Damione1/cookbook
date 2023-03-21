package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	database "github.com/Damione1/portfolio-playground/db"
	"github.com/Damione1/portfolio-playground/pkg/grpcApi"
	"github.com/Damione1/portfolio-playground/pkg/pb"
	"github.com/Damione1/portfolio-playground/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config. %v", err)
	}

	db, err := database.ConnectDb(config)
	if err != nil {
		log.Fatalf("Failed to connect to database. %v", err)
	}

	runGrpcServer(config, db)
}

func runGrpcServer(config util.Config, store *sql.DB) {
	server, err := grpcApi.NewServer(config)
	if err != nil {
		fmt.Errorf("Failed to create gRPC server. %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPortfolioServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		fmt.Errorf("Failed to listen. %v", err)
	}

	fmt.Printf("Starting gRPC server on port %s", listener.Addr().String())
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Errorf("Failed to serve gRPC server over port %s. %v", listener.Addr().String(), err)
	}
}
