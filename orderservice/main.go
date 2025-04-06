package main

import (
	"context"
	"log"
	"net"

	pb "github.com/122140015devaahmad/grpc-saga/proto/order"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedOrderServiceServer
}

func (s *server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	log.Printf("Order created: %s", req.OrderId)
	return &pb.CreateOrderResponse{
		OrderId: req.OrderId,
		Status:  "PENDING",
	}, nil
}

func (s *server) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	log.Printf("Order cancelled: %s", req.OrderId)
	return &pb.CancelOrderResponse{
		Message: "Order cancelled",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &server{})
	log.Printf("Order service listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
