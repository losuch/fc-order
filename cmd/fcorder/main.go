package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"github.com/losuch/fc-order/api"
	db "github.com/losuch/fc-order/db/sqlc"
	"github.com/losuch/fc-order/pb"
	"github.com/losuch/fc-order/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
    log.Printf("Received request for method: %s", info.FullMethod)
    return handler(ctx, req)
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        startTime := time.Now()

        next.ServeHTTP(w, r)

        log.Printf(
            "%s %s %s %s",
            r.Method,
            r.RequestURI,
            r.RemoteAddr,
            time.Since(startTime),
        )
    })
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	err = conn.Ping()
    if err != nil {
        log.Fatal("cannot ping db:", err)
    }

	store := db.NewStore(conn)
	go runGatewayServer(config, store)
	runGrpcServer(config, store)
}

func runGatewayServer(config util.Config, store db.Store) {

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterFcOrderHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register service:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", loggingMiddleware(grpcMux))

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("start HTTP server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot start HTTP gateway server:", err)
	}
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor))
	pb.RegisterFcOrderServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server:",err)
	}
}

