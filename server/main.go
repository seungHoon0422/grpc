package main

import (
	"context"
	"log"
	"net"

	pb "github.com/seungHoon0422/grpc/proto"

	"google.golang.org/grpc"
)

// server는 MyServiceServer 인터페이스를 구현합니다.
type server struct {
	pb.UnimplementedMyServiceServer
}

// Reverse 메서드는 문자열을 뒤집는 RPC를 처리합니다.
func (s *server) Reverse(ctx context.Context, req *pb.ReverseRequest) (*pb.ReverseResponse, error) {
	input := req.GetInput()
	reversed := reverseString(input)
	response := &pb.ReverseResponse{
		Output: reversed,
	}
	return response, nil
}

// CheckEvenOdd 메서드는 숫자가 홀수인지 짝수인지 확인하는 RPC를 처리합니다.
func (s *server) CheckEvenOdd(ctx context.Context, req *pb.CheckEvenOddRequest) (*pb.CheckEvenOddResponse, error) {
	number := req.GetNumber()
	isEven := number%2 == 0
	response := &pb.CheckEvenOddResponse{
		IsEven: isEven,
	}
	return response, nil
}

// reverseString 함수는 문자열을 뒤집습니다.
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// gRPC 서버를 50051 포트에서 시작합니다.
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// gRPC 서버 인스턴스 생성합니다.
	s := grpc.NewServer()

	// MyService 서비스를 등록합니다.
	pb.RegisterMyServiceServer(s, &server{})

	log.Println("Server started on port 50051...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
