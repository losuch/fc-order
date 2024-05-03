package api

import (
	"context"

	"github.com/losuch/fc-order/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateAccount creates a new account
func (server *Server) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}