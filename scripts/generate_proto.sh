#!/bin/sh

protoc --go-grpc_out=pkg/pb --go_out=pkg/pb --proto_path=proto --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative --grpc-gateway_out=pkg/pb --grpc-gateway_opt=paths=source_relative ./proto/*.proto
