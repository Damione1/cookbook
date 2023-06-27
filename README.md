# Damien's Cookbook

Damien's Cookbook is a personal project aimed at learning various technologies and best practices in the field of software development. This project focuses on creating a cookbook platform where users can register, submit their recipes with pictures, instructions, and the list of ingredients. The backend is developed using Golang, and the frontend will be developed using Nuxt.js in the second phase. Most of the logic and validation will be handled on the backend to ensure data integrity and security (and also because it's the funniest part to develop).

The main technologies and concepts involved in this project are:

- [x] Protocol Buffers & gRPC
- [x] RESTful API (With GRPC Gateway)
- [x] SQLBoiler as ORM
- [x] SQL Migrations (with Golang Migrate)
- [x] JWT authentication (with PASETO)
- [x] Docker Compose
- [-] Strict validations
- [-] Strict Linter
- [] Roles & Permissions
- [x] Tilt (for hot reloading)
- [x] Automated documentation (with Swagger)
- [-] Logging
- [] Media management
- [] Redis caching
- [] Email sending
- [] Best practices
- [] Testing
- [] Automated deployments
- [] Google Cloud Services

The list will be updated as the project progresses. The goal is to learn to use most of the microservices best practices and technologies.

## Getting Started

To start the project with tilt, run the following command:

```bash
tilt up
```
Then to access the tilt UI, go to the following URL:

```bash
http://localhost:10350/
```


To connect to the PostgreSQL database, use the following command:

```bash
docker-compose exec db psql -U postgres -d cookbook
```

To access the Swagger UI, go to the following URL once the project is running:

```bash
http://localhost:8080/swagger/
```

## Project Overview (Current State)

The cookbook platform allows multiple users to register and submit their recipes with pictures, steps, and the list of ingredients. Users can view the ingredients in multiple units and adjust the quantities according to the desired serving size. They can also add multiple recipes to a cart (similar to an e-commerce platform) and get a comprehensive list of necessary ingredients. Identical ingredients will be merged, and the quantities will be added accordingly.

## Future Development

The second phase of this project will involve developing the frontend using Nuxt.js to provide a seamless user experience. Furthermore, additional features and improvements will be made to the backend to enhance the overall functionality of the platform.

Stay tuned for updates and new developments!
