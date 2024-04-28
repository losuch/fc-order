package api

import (
	"context"

	"github.com/losuch/fc-order/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateAccount creates a new account
func (s *Server) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	response := &pb.CreateAccountResponse{
		Account: &pb.Account{
			Email: req.Email,
			Role:  req.Role,
			AcccountId: 1,
			CreatedAt: &timestamppb.Timestamp{
				Seconds: 1630483200,
			},
		},
	}
	return response, nil
}