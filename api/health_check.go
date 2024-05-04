package api

import (
	"context"

	"github.com/losuch/fc-order/pb" // replace with your actual import path
)

// Check implements the HealthCheck service
func (server *Server) Check(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {

    return &pb.HealthCheckResponse{Status: "OK"}, nil
}
