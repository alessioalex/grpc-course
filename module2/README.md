## .

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/hello.proto

https://github.com/fullstorydev/grpcurl

```
$ grpcurl -d '{"name": "Chris" }' \
        -import-path ./proto -proto hello.proto \
        -plaintext localhost:50051 \
        hello.HelloService/SayHello
{
  "message": "Hello Chris!"
}
```

gRPC status codes - https://grpc.io/docs/guides/status-codes/
Using gRPC status codes - https://www.bytesizego.com/blog/grpc-status-codes
