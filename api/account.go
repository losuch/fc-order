package api

import (
	"context"

	db "github.com/losuch/fc-order/db/sqlc"
	"github.com/losuch/fc-order/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateAccount creates a new account
func (server *Server) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.Account, error) {

	arg := db.CreateAccountParams{

		Role: req.Role,
		HashedPassword: req.Password,
		Email:          req.Email,
	}

	account, err := server.store.CreateAccount(ctx, arg)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)	
		
	}
		
	return &pb.Account{
			AccountId: account.AccountID,
			Email: account.Email,
			Role: account.Role,
			CreatedAt: timestamppb.New(account.CreatedAt.Time),
	}, nil
	
}

// GetAccountList returns a list of accounts
func (server *Server) GetAccountList(ctx context.Context, req *pb.GetAccountListRequest) (*pb.GetAccountListResponse, error) {
	
	accounts, err := server.store.GetAccountList(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %v", err)
	}
	
	var pbAccounts []*pb.Account
	for _, account := range accounts {
		pbAccount := &pb.Account{
			AccountId: account.AccountID,
			Email: account.Email,
			Role: account.Role,
			CreatedAt: timestamppb.New(account.CreatedAt.Time),
		}
		pbAccounts = append(pbAccounts, pbAccount)
	}
	
	return &pb.GetAccountListResponse{
		Accounts: pbAccounts,
	}, nil
}