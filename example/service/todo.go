package main

import (
	"context"
	todopb "example/gen"
	"fmt"

	"github.com/jaevor/go-nanoid"
	"go.einride.tech/aip/fieldmask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func newTodoService() *todoService {
	// ignore error since len is known to be correct
	gen, _ := nanoid.Standard(21)
	return &todoService{
		db:    map[string]*todopb.Todo{},
		genId: gen,
	}
}

type todoService struct {
	todopb.UnimplementedTodoServiceServer
	db    map[string]*todopb.Todo
	genId func() string
}

func (s *todoService) ListTodos(ctx context.Context, in *emptypb.Empty) (*todopb.Todos, error) {
	ret := &todopb.Todos{}
	for _, v := range s.db {
		ret.Items = append(ret.Items, v)
	}
	return ret, nil
}

func (s *todoService) GetTodo(ctx context.Context, in *todopb.TodoId) (*todopb.Todo, error) {
	item, ok := s.db[in.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("todo with id %s not found", in.Id))
	}
	return item, nil
}

func (s *todoService) CreateTodo(ctx context.Context, in *todopb.Todo) (*todopb.Todo, error) {
	if in.Id == "" {
		in.Id = s.genId()
	}
	if _, ok := s.db[in.Id]; ok {
		return nil, status.Errorf(codes.AlreadyExists, fmt.Sprintf("todo with id %s already exists", in.Id))
	}
	s.db[in.Id] = in
	return in, nil
}

func (s *todoService) UpdateTodo(ctx context.Context, in *todopb.UpdateTodoRequest) (*todopb.Todo, error) {
	cur, ok := s.db[in.Todo.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("todo with id %s not found", in.Todo.Id))
	}
	fieldmask.Update(in.UpdateMask, cur, in.Todo)
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}

func (s *todoService) DeleteTodo(ctx context.Context, in *todopb.TodoId) (*todopb.TodoId, error) {
	_, ok := s.db[in.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("todo with id %s not found", in.Todo.Id))
	}
	delete(s.db, in.Id)
	return in, nil
}
