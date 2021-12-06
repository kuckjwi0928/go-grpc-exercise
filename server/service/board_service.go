package service

import (
	"context"
	"encoding/json"
	. "go-grpc-exercise/gen/proto/board"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var mock BoardsMock

type BoardsMock struct {
	Boards []BoardMock `json:"boards"`
}

type BoardMock struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
}

func init() {
	f, err := os.Open("./server/mocking/mock.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	b, _ := ioutil.ReadAll(f)
	err = json.Unmarshal(b, &mock)
	if err != nil {
		log.Fatal(err)
	}
}

type BoardService struct {
	UnimplementedBoardServer
}

func (b *BoardService) GetBoard(ctx context.Context, req *BoardRequest) (*BoardResponse, error) {
	for _, board := range mock.Boards {
		if board.Id == req.Id {
			return &BoardResponse{
				Id:        board.Id,
				Title:     board.Title,
				Content:   board.Content,
				CreatedAt: timestamppb.New(time.Unix(board.CreatedAt, 0)),
			}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Board not found")
}
