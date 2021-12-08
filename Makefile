init:
	GO111MODULE=off go install google.golang.org/protobuf/cmd/protoc-gen-go
	go mod download

generate: init
	go generate ./...