pb "github.com/seungHoon0422/grpc/proto"




protoc -I=. \
            --go_out . --go_opt paths=source_relative \
            --go-grpc_out . --go-grpc_opt paths=source_relative \
            proto/myservice.proto


### Server 실행

server 파일의 위치로 이동
cd module/server
server 실행 명령어 입력

go run main.go


### client 실행

client 파일의 위치로 이동
cd module/client
client 실행 명령어 입력

go run client.go


# Golang과 gRPC를 활용한 서버 구현

<aside>
💡 1. 문자열을 입력 받은뒤 뒤집어서 출력해 주는 RPC
2. 정수를 입력받아서 홀수, 짝수 여부를 알려주는 RPC

</aside>