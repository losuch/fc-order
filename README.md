# fc-order

filip-club order backen API

## golang-migrate

Tool to maintain database schema migrations.
https://github.com/golang-migrate/migrate

### Install

`$ brew install golang-migrate`

### Create Migration file

`$ migrate create -ext sql -dir db/migration -seq init_schema`

### Migrate Up

`$ migrate -path db/migration -database "postgres://localhost:5432/filipclub?sslmode=disable" up`

## Database

Database schema design: https://dbdiagram.io/

### SQLC

SQLC is a code generation tool for writing SQL queries in Go. It is designed to replace many of the "ORM" style libraries and provide a much simpler and more performant interface to your database of choice. The sqlc.yaml contains the configuration for the sqlc tool.

### gRPC

Use the evans tool to interact with the gRPC server.

`$ evans -r repl --host localhost --port 9090`

#### grpc-gateway

The grpc-gateway is a plugin of the Google protocol buffers compiler (protoc). It reads gRPC service definitions and generates a reverse-proxy server which translates a RESTful JSON API into gRPC. This server is generated according to the google.api.http annotations in your service definitions.

#### Usage

To use the grpc-gateway, you need to follow these steps:

1. Install the tools by adding toold.go:

   ```
   // +build tools

   package tools

   import (
       _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
       _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
       _ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
       _ "google.golang.org/protobuf/cmd/protoc-gen-go"
   )
   ```

   and run the following command:
   `go mod tidy`

2. Install the necessary tools by running the following command:

   ```
   $ go install \
       github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
       github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
       google.golang.org/protobuf/cmd/protoc-gen-go \
       google.golang.org/grpc/cmd/protoc-gen-go-grpc
   ```

3. Implement your gRPC service and add the necessary annotations for the gateway:

   ```protobuf
   service YourService {
     rpc YourMethod (YourRequest) returns (YourResponse) {
        option (google.api.http) = {
           post: "/v1/yourmethod"
           body: "*"
        };
     }
   }
   ```

4. Build and run your gateway server:

   ```go
   package main

   import (
     "context"
     "log"
     "net/http"

     "github.com/grpc-ecosystem/grpc-gateway/runtime"
     "google.golang.org/grpc"
   )

   func main() {
     ctx := context.Background()
     ctx, cancel := context.WithCancel(ctx)
     defer cancel()

     mux := runtime.NewServeMux()
     opts := []grpc.DialOption{grpc.WithInsecure()}

     err := yourpb.RegisterYourServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
     if err != nil {
        log.Fatalf("failed to register gateway: %v", err)
     }

     http.ListenAndServe(":8080", mux)
   }
   ```
