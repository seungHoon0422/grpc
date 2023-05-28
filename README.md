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
