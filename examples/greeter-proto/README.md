# grpc

1. install protoc
$ brew update
$ brew install protobuf

2. install go dependecies
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

3. update PATH

- show GOPATH
$ go env GOPATH

- vim ~/.zshrc
export PATH=$PATH:/path/to/your/go/bin

4. protoc generation

$ protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       greeter.proto

$ go mod tidy