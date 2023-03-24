# portfolio-playground

## To start the project, run the following command:

```bash
docker-compose up
```

## To start the project with hot reloading, run the following command:

```bash
docker-compose up web-with-air
```

## To generate models from the database, run the following command:

```bash
docker-compose up generate-models
```

## To generate the protobuf files, run the following command:

```bash
protoc --go-grpc_out=pkg/pb --go_out=pkg/pb --proto_path=proto --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative --grpc-gateway_out=pkg/pb --grpc-gateway_opt=paths=source_relative ./proto/*.proto
```
**Note: You need to have the `protoc` binary installed on your machine. You can download it from [here](https://github.com/protocolbuffers/protobuf/releases/), copy it in a folder and put the path of his `bin` folder in your `PATH` environment variable.**

To connect to the postgres database, you can use the following command:

```bash
docker-compose exec db psql -U postgres -d portfolio_playground
```

To test the gRPC server, you can use the following command:

```bash
evans --host localhost --port 9090 -r repl
```
Then, you can list the available services with the `service list` command and call the `GetUser` method with the `call GetUser` command.
