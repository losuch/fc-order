package api

import (
	"context"
	"log"
	"time"

	"github.com/losuch/fc-order/pb" // replace with your actual import path
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Check implements the HealthCheck service returns a status and timestamp
func (server *Server) Check(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {

	// log the request path
	log.Printf("Received: %v", in)
		
    return &pb.HealthCheckResponse{Status: "OK", Timestamp: timestamppb.New(time.Now())},
		nil
}
