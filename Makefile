.PHONY:hello
hello:
	@echo "Hello"

world: hello
	@echo "World"	

compile:
	protoc -I=. --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative protos/v1/user/user.proto

run:
	@go run server/server.go

build:
	@go build -o bin/ app/user/server/server.go
	@go build -o bin/ app/user/client/client.go

.PHONY:clean
clean:
	@rm bin/* 

test:
	@go test -v

all: hello run build test

cc: 
	@echo "Cross Compile"
	set GOOS=darwin& set GOARCH=386& go build -o bin/server-darwin-386.exe server/server.go
	set GOOS=linux& set GOARCH=amd64& go build -o bin/server-linux-amd64.exe server/server.go
	set GOOS=windows& set GOARCH=arm& go build -o bin/server-windows-arm.exe server/server.go

cclinux:
	@echo "Cross Compile"
	GOOS=darwin GOARCH=386 go build -o bin/hello-darwin-386.exe server/server.go
	GOOS=linux GOARCH=amd64 go build -o bin/hello-linux-amd64.exe server/server.go
	GOOS=windows GOARCH=arm go build -o bin/hello-windows-arm.exe server/server.go

