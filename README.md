
# Golang과 gRPC를 활용한 서버 구현

</br>

<aside>
💡 1. 문자열을 입력 받은뒤 뒤집어서 출력해 주는 RPC</br>
💡 2. 정수를 입력받아서 홀수, 짝수 여부를 알려주는 RPC
</aside>

</br>


### 프로젝트 구조

</br>

```
|— MODULE (module name : github.com/seungHoon0422/grpc/proto)
    |— client
        |—client.go

    |— server
        |—main.go

    |— proto
        |— myservice.proto
        |— myservice.pb.go
        |—myservice_grpc.pb.go

    |— go.mod
    |— go.sum
```

</br>

### myservice.proto 파일 컴파일

1. proto 디렉토리로 이동
2. myservice.proto 파일 컴파일 진행
```bash
protoc -I=. \
            --go_out . --go_opt paths=source_relative \
            --go-grpc_out . --go-grpc_opt paths=source_relative \
            proto/myservice.proto
```
3. 컴파일을 진행하면 `myservice.pb.go`, `myservice_grpc.pb.go` 파일이 생성됩니다.

### server 실행

1. server 디렉토리로 이동
2. server 실행
3. 서버는 localhost:50051 port에서 실행 
```bash
go run main.go
```


### client 실행

1. client 디렉토리로 이동
2. client 서비스 실행
```bash
go run client.go
```
3. 



- github source code : https://github.com/seungHoon0422/grpc


