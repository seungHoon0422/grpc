package main

import (
	"context"
	"log"
	"net"

	pb "github.com/seungHoon0422/grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
	// pb.UnimplementedMyServiceServer
}

func (s *server) Reverse(ctx context.Context, req *pb.ReverseRequest) (*pb.ReverseResponse, error) {
	input := req.GetInput()
	reversed := reverseString(input)
	response := &pb.ReverseResponse{
		Output: reversed,
	}
	return response, nil
}

func (s *server) CheckEvenOdd(ctx context.Context, req *pb.CheckEvenOddRequest) (*pb.CheckEvenOddResponse, error) {
	number := req.GetNumber()
	isEven := number%2 == 0
	response := &pb.CheckEvenOddResponse{
		IsEven: isEven,
	}
	return response, nil
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// pb.RegisterMyServiceServer(s, &server{})

	log.Println("Server started on port 50051...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
