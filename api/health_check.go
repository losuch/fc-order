package api

import (
	"context"
	"time"

	"github.com/losuch/fc-order/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Check implements the HealthCheck service returns a status and timestamp
func (server *Server) Check(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
		
    return &pb.HealthCheckResponse{Status: "OK", Timestamp: timestamppb.New(time.Now())},
		nil
}
