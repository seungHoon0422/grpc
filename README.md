
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
|— MODULE (module name : github.com/seungHoon0422/grpc)

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

---

</br>


### 프로젝트를 시작하기 위한 사전 준비

</br>

- Go 언어 설치
Go 언어를 사용하여 gRPC 서버를 작성하기 위해서는 먼저 Go 언어를 설치해야 합니다. 

</br>

- 프로토콜 버퍼 설치: gRPC는 프로토콜 버퍼(Protocol Buffers)를 사용하여 데이터를 직렬화하고 RPC 인터페이스를 정의합니다. 프로토콜 버퍼 컴파일러인 protoc를 설치해야 합니다. 

</br>

- Go 언어용 프로토콜 버퍼 플러그인을 설치해야 합니다. 터미널에서 다음 명령을 실행하여 설치합니다.

```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

</br>

---

</br>

### myservice.proto 파일 컴파일

</br>

1. proto 디렉토리로 이동
2. myservice.proto 파일 컴파일 진행
```bash
$ protoc -I=. \
            --go_out . --go_opt paths=source_relative \
            --go-grpc_out . --go-grpc_opt paths=source_relative \
            proto/myservice.proto
```
3. 컴파일을 진행하면 `myservice.pb.go`, `myservice_grpc.pb.go` 파일이 생성됩니다.


</br>

---

</br>


### server 실행

</br>

1. server 디렉토리로 이동
2. server 실행
3. 서버는 localhost:50051 port에서 실행 
```bash
$ go run main.go
```

</br>

---

</br>


### client 실행

</br>

1. client 디렉토리로 이동
2. client 서비스 실행
```bash
$ go run client.go
```

3. client 코드 진행
    - 3.1> string을 입력받아 reverse된 string을 보여줍니다.
    - 3.2> 정수를 입력받아 홀수(odd),짝수(even) 여부를 알려줍니다.

</br>


- github source code : https://github.com/seungHoon0422/grpc


