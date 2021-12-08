package server

import (
	"context"
	pb "go-grpc-exercise/gen/proto/board"
	"go-grpc-exercise/server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterBoardServer(s, &service.BoardService{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func TestGetBoard(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "test", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial test: %v", err)
	}
	defer conn.Close()
	client := pb.NewBoardClient(conn)
	res, err := client.GetBoard(ctx, &pb.BoardRequest{Id: 1})
	if err != nil {
		t.Fatalf("GetBoard failed: %v", err)
	}
	log.Printf("Response: %+v", res)
}
