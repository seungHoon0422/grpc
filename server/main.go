package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "module/proto" // "your-module-name"을 프로토콜 버퍼 파일이 위치하는 모듈 이름으로 변경

	"module/evenoddpkg" // 홀짝 여부 확인 로직을 포함한 패키지 이름으로 변경
	"module/reversepkg" // 문자열 뒤집기 로직을 포함한 패키지 이름으로 변경
)

type server struct{}

func (s *server) Reverse(ctx context.Context, req *pb.ReverseRequest) (*pb.ReverseResponse, error) {
	input := req.GetInput()
	reversed := reversepkg.ReverseString(input)

	return &pb.ReverseResponse{Output: reversed}, nil
}

func (s *server) CheckEvenOdd(ctx context.Context, req *pb.CheckEvenOddRequest) (*pb.CheckEvenOddResponse, error) {
	number := req.GetNumber()
	isEven := evenoddpkg.CheckEvenOdd(number)

	return &pb.CheckEvenOddResponse{IsEven: isEven}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})

	fmt.Println("gRPC server is running...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
