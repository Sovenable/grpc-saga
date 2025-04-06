package main

import (
	"context"
	"log"
	"net"

	pb "github.com/122140015devaahmad/grpc-saga/proto/shipping"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedShippingServiceServer
}

func (s *server) StartShipping(ctx context.Context, req *pb.StartShippingRequest) (*pb.StartShippingResponse, error) {
	log.Printf("Shipping started for order: %s", req.OrderId)
	return &pb.StartShippingResponse{
		OrderId: req.OrderId,
		Status:  "SHIPPED",
	}, nil
}

func (s *server) CancelShipping(ctx context.Context, req *pb.CancelShippingRequest) (*pb.CancelShippingResponse, error) {
	log.Printf("Shipping cancelled for order: %s", req.OrderId)
	return &pb.CancelShippingResponse{
		Message: "Shipping cancelled",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterShippingServiceServer(s, &server{})
	log.Printf("Shipping service listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
