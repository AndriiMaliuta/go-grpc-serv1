gen:
	protoc --proto_path=proto proto/*.proto --go_out=server --go-grpc_out=server
	protoc --proto_path=proto proto/*.proto --go_out=client --go-grpc_out=client

#   windows
gen_win:
	protoc --proto_path=. --go_out=. ./proto/messages.proto

get_plugin:
	protoc blog/blogpb/blog.proto --go-out=plugins=grpc:.

linux_clean:
	rm -rf server/pb/
	rm -rf client/pb/

server:
	go run server/main.go

linux_install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	brew install protobuf
	brew install clang-format
	brew install grpcurl
	export PATH=$PATH:$(go env GOPATH)/bin

test:
	rm -rf tmp && mkdir tmp
	go test -cover -race serializer/*.go