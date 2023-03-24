package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	database "github.com/Damione1/portfolio-playground/db"
	"github.com/Damione1/portfolio-playground/pkg/grpcApi"
	"github.com/Damione1/portfolio-playground/pkg/pb"
	"github.com/Damione1/portfolio-playground/util"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Println(fmt.Printf("Failed to load config. %v", err))
	}

	db, err := database.ConnectDb(&config)
	if err != nil {
		log.Println(fmt.Printf("Failed to connect to database. %v", err))
	}

	go runGatewayServer(config, db)
	runGrpcServer(config, db)

}

func runGrpcServer(config util.Config, store *sql.DB) {
	log.Println("Starting gRPC server")
	server, err := grpcApi.NewServer(config)
	if err != nil {
		log.Println(fmt.Printf("Failed to create gRPC server. %v", err))
	}
	log.Println("gRPC server created")
	grpcServer := grpc.NewServer()
	pb.RegisterPortfolioServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	log.Println("Starting to listen on port " + config.GRPCServerAddress)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Println(fmt.Printf("Failed to listen. %v", err))
	}

	log.Println(fmt.Printf("Starting gRPC server on port %s.", listener.Addr().String()))
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(fmt.Printf("Failed to serve gRPC server over port %s. %v", listener.Addr().String(), err))
	}
}

func runGatewayServer(config util.Config, store *sql.DB) {
	log.Println("Starting gRPC server")
	server, err := grpcApi.NewServer(config)
	if err != nil {
		log.Println(fmt.Printf("Failed to create gRPC server. %v", err))
	}

	grpcMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterPortfolioServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal(fmt.Printf("Failed to register HTTP gateway server. %v", err))
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal(fmt.Printf("Failed to listen. %v", err))
	}

	log.Println(fmt.Printf("Starting HTTP gateway server on port %s.", listener.Addr().String()))

	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal(fmt.Printf("Failed to serve HTTP gateway server over port %s. %v", listener.Addr().String(), err))
	}
}
