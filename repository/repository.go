package repository

import (
	"context"

	"trixtaro.dev/go/grpc/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) (error)
	GetTest(ctx context.Context, id string) (*models.Test, error)
	SetTest(ctx context.Context, test *models.Test) (error)
	SetQuestion(ctx context.Context, question *models.Question) (error)
}

var implementation Repository

func SetImplementation(repo Repository) {
	implementation = repo
}

func SetStudent(ctx context.Context, student *models.Student) (error) {
	return implementation.SetStudent(ctx, student)
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}

func SetTest(ctx context.Context, test *models.Test) (error) {
	return implementation.SetTest(ctx, test)
}

func SetQuestion(ctx context.Context, question *models.Question) (error) {
	return implementation.SetQuestion(ctx, question)
}
