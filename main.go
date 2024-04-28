package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/losuch/fc-order/api"
	db "github.com/losuch/fc-order/db/sqlc"
	"github.com/losuch/fc-order/pb"
	"github.com/losuch/fc-order/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

func main() {

	// loading env file with config
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFcOrderServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatal("cannot listen:", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot server:", err)
	}	
}
