package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/israelmiranda/go-samples/examples/coupon-proto"
)

type server struct {
	pb.UnimplementedCouponServer
}

func (s *server) GetCoupons(ctx context.Context, empty *emptypb.Empty) (*pb.CouponsListResponse, error) {
	coupons := []*pb.CouponResponse{
		{Code: "GRPC-123"},
		{Code: "PROTO-456"},
		{Code: "GO-789"},
	}

	couponsList := &pb.CouponsListResponse{
		Coupons: coupons,
	}

	return couponsList, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCouponServer(s, &server{})
	fmt.Println("Server listening on port 50051...")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
