package server

import (
	"context"
	"fmt"

	"trixtaro.dev/go/grpc/models"
	"trixtaro.dev/go/grpc/repository"
	"trixtaro.dev/go/grpc/studentpb"
)

type Serve struct {
	repo repository.Repository
	studentpb.UnimplementedStudentServiceServer
}

func NewStudentServer(repo repository.Repository) *Serve {
	return &Serve{repo: repo}
}

func (s *Serve) GetStudent(context context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repo.GetStudent(context, req.Id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &studentpb.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *Serve) SetStudent(context context.Context, req *studentpb.Student) (*studentpb.SetStudentResponse, error) {
	student := &models.Student{
		Id:   req.GetId(),
		Name: req.GetName(),
		Age:  req.GetAge(),
	}

	err := s.repo.SetStudent(context, student)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &studentpb.SetStudentResponse{
		Id: student.Id,
	}, nil
}