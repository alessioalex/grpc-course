## .

https://grpc.io/docs/protoc-installation/
https://grpc.io/docs/languages/go/quickstart/

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

export PATH="$PATH:$(go env GOPATH)/bin"


protoc --go_out=. --go_opt=paths=source_relative proto/hello.proto
go mod tidy
go mod vendor
