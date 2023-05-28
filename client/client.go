package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/seungHoon0422/grpc/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// gRPC 서버에 연결
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// gRPC 클라이언트 생성
	client := pb.NewMyServiceClient(conn)

	// 사용자로부터 문자열 입력 받기
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string to reverse: ")
	input, _ := reader.ReadString('\n')

	// 문자열 뒤집기 RPC 호출
	reverseReq := &pb.ReverseRequest{
		Input: input,
	}
	reverseRes, err := client.Reverse(context.Background(), reverseReq)
	if err != nil {
		log.Fatalf("Reverse RPC failed: %v", err)
	}
	fmt.Printf("Reversed output: %s\n", reverseRes.GetOutput())

	// 사용자로부터 정수 입력 받기
	fmt.Print("Enter an integer to check even/odd: ")
	var number int32
	_, err = fmt.Scanf("%d", &number)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	// 홀수, 짝수 확인 RPC 호출
	checkReq := &pb.CheckEvenOddRequest{
		Number: number,
	}
	checkRes, err := client.CheckEvenOdd(context.Background(), checkReq)
	if err != nil {
		log.Fatalf("CheckEvenOdd RPC failed: %v", err)
	}
	result := "even"
	if !checkRes.GetIsEven() {
		result = "odd"
	}
	fmt.Printf("Number is %s\n", result)
}
