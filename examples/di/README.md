# wire

1. install go dependecies
$ go install github.com/google/wire/cmd/wire@latest

2. update PATH

- show GOPATH
$ go env GOPATH

- vim ~/.zshrc
export PATH=$PATH:/path/to/your/go/bin

3. create wire.go
//go:build wireinject
// +build wireinject

put those lines before package name

$ go mod tidy

4. wire generation
$ wire

5. run
$ go run main.go wire_gen.go