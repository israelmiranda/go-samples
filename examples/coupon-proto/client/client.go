package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/israelmiranda/go-samples/examples/coupon-proto"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCouponClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetCoupons(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not get coupons: %v", err)
	}

	fmt.Println("Coupons received:")
	for _, coupon := range r.GetCoupons() {
		fmt.Println("-", coupon.GetCode())
	}
}
