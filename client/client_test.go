package client

import (
	"testing"
)

var client *BoardCaller

func init() {
	client = &BoardCaller{}
}

func TestBoardCaller_CallGetBoard(t *testing.T) {
	board, err := client.CallGetBoard(1)
	if err != nil {
		t.Error(err)
	}
	if board.GetId() != 1 {
		t.Error("expected 1")
	}
	if board.GetTitle() != "Hi kuckjwi~" {
		t.Error("expected Hi kuckjwi~")
	}
	if board.GetCreatedAt().GetSeconds() != 1638798195 {
		t.Error("expected 1638798195")
	}
	_, err = client.CallGetBoard(3)
	if err == nil {
		t.Error("expected error")
	}
}
