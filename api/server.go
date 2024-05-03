package api

import (
	db "github.com/losuch/fc-order/db/sqlc"
	"github.com/losuch/fc-order/pb"

	"github.com/losuch/fc-order/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	pb.UnimplementedFcOrderServer
	config     util.Config
	store      db.Store
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	

	server := &Server{
		config:     config,
		store:      store,
	}

	return server, nil
}
