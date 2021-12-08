package gen

//go:generate -command compile_proto protoc -I../ --go_out=../gen --go_opt=paths=source_relative --go-grpc_out=../gen --go-grpc_opt=paths=source_relative
//go:generate compile_proto ../proto/board/board.proto
