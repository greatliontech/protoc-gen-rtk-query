package main

import (
	"context"
	"testing"

	todopb "github.com/greatliontech/protoc-gen-rtk-query/example/service/gen"
	"github.com/greatliontech/protoc-gen-rtk-query/example/service/mocks"
	"go.uber.org/mock/gomock"
)

func TestTodoService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoService := mocks.NewMockTodoServiceServer(ctrl)

	mockTodoService.EXPECT().CreateTodo(gomock.Any(), &todopb.Todo{
		Id:    "testid",
		Title: "Buy milk",
	}).Return(&todopb.Todo{
		Id:    "testid",
		Title: "Buy milk",
		State: todopb.State_TODO,
	}, nil)

	s := newTodoService()

	res, err := s.CreateTodo(context.Background(), &todopb.Todo{
		Id:    "testid",
		Title: "Buy milk",
	})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if res.Id != "testid" {
		t.Errorf("Expected: testid, got: %v", res.Id)
	}
}
