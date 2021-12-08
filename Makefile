#generate:
#	protoc -I. --go_out=gen --go_opt=paths=source_relative \
#	--go-grpc_out=gen --go-grpc_opt=paths=source_relative proto/**/*.proto

init:
	GO111MODULE=off go install google.golang.org/protobuf/cmd/protoc-gen-go
	go mod download

generate: init
	go generate ./...