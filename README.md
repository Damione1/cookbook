# Portfolio Playground

Portfolio Playground is a personal project aimed at learning various technologies and best practices in the field of software development. This project focuses on creating a cookbook platform where users can register, submit their recipes with pictures, steps, and the list of ingredients. The backend is developed using Golang, and the frontend will be developed using Nuxt.js in the second phase.

The main technologies and concepts involved in this project are:

- Protocol Buffers & gRPC
- SQLBoiler
- JWT authentication (with PASETO)
- Docker Compose
- Tilt
- Email sending
- Best practices
- Testing
- Automated deployments
- Google Cloud Services

## Getting Started

To start the project with tilt, run the following command:

```bash
tilt up
```

To start the project using Docker Compose, run the following command:

```bash
docker-compose up
```

For hot reloading, start the project with the following command:

```bash
docker-compose up web-with-air
```

To generate models from the database, run the following command:

```bash
docker-compose up generate-models
```

To generate the protobuf files, run the following command:

```bash
protoc --go-grpc_out=pkg/pb --go_out=pkg/pb --proto_path=proto --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative --grpc-gateway_out=pkg/pb --grpc-gateway_opt=paths=source_relative ./proto/*.proto
```

**Note: You need to have the `protoc` binary installed on your machine. You can download it from [here](https://github.com/protocolbuffers/protobuf/releases/), copy it in a folder and put the path of its `bin` folder in your `PATH` environment variable.**

To connect to the PostgreSQL database, use the following command:

```bash
docker-compose exec db psql -U postgres -d portfolio_playground
```

To test the gRPC server, use the following command:

```bash
evans --host localhost --port 9090 -r repl
```

Then, you can list the available services with the `service list` command and call the `GetUser` method with the `call GetUser` command.

## Project Overview

The cookbook platform allows multiple users to register and submit their recipes with pictures, steps, and the list of ingredients. Users can view the ingredients in multiple units and adjust the quantities according to the desired serving size. They can also add multiple recipes to a cart (similar to an e-commerce platform) and get a comprehensive list of necessary ingredients. Identical ingredients will be merged, and the quantities will be added accordingly.

## Future Development

The second phase of this project will involve developing the frontend using Nuxt.js to provide a seamless user experience. Furthermore, additional features and improvements will be made to the backend to enhance the overall functionality of the platform.

Stay tuned for updates and new developments!
