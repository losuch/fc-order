package api

import (
	db "github.com/losuch/fc-order/db/sqlc"
	"github.com/losuch/fc-order/pb"

	// "phonebook/token"
	// "phonebook/util"

	"github.com/losuch/fc-order/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	pb.UnimplementedFcOrderServer
	config     util.Config
	store      db.Store
	// tokenMaker token.Maker
	
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	// tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot create token maker: %w", err)
	// }
	server := &Server{
		config:     config,
		store:      store,
		// tokenMaker: tokenMaker,
	}

	return server, nil
}


