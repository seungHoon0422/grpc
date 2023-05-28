package main

import (
	"context"
	"log"

	pb "github.com/seungHoon0422/grpc/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMyServiceClient(conn)

	// Reverse RPC
	reverseReq := &pb.ReverseRequest{
		Input: "Hello, World!",
	}
	reverseRes, err := client.Reverse(context.Background(), reverseReq)
	if err != nil {
		log.Fatalf("Reverse RPC failed: %v", err)
	}
	log.Printf("Reversed output: %s", reverseRes.GetOutput())

	// CheckEvenOdd RPC
	checkReq := &pb.CheckEvenOddRequest{
		Number: 7,
	}
	checkRes, err := client.CheckEvenOdd(context.Background(), checkReq)
	if err != nil {
		log.Fatalf("CheckEvenOdd RPC failed: %v", err)
	}
	result := "even"
	if !checkRes.GetIsEven() {
		result = "odd"
	}
	log.Printf("Number is %s", result)
}
