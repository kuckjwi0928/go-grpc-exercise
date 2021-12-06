package client

import (
	"context"
	pb "go-grpc-exercise/gen/proto/board"
	"google.golang.org/grpc"
	"time"
)

type BoardCaller struct{}

func (b *BoardCaller) CallGetBoard(id int64) (*pb.BoardResponse, error) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewBoardClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.GetBoard(ctx, &pb.BoardRequest{Id: id})
	return res, err
}
