package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"trixtaro.dev/go/grpc/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	fmt.Println("Inserting student...")
	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO students (id, name, age) VALUES ($1, $2, $3)",
		student.Id, student.Name, student.Age,
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (repo *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func ()  {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var student = models.Student{}
	for rows.Next() {
		err = rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}
		return &student, nil
	}
	return &student, nil
}

func (repo *PostgresRepository) SetTest(ctx context.Context, test *models.Test) error {
	fmt.Println("Inserting test...")
	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO tests (id, name) VALUES ($1, $2)",
		test.Id, test.Name,
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (repo *PostgresRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name FROM tests WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func ()  {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var test = models.Test{}
	for rows.Next() {
		err = rows.Scan(&test.Id, &test.Name)
		if err != nil {
			return nil, err
		}
		return &test, nil
	}
	return &test, nil
}

func (repo *PostgresRepository) SetQuestion(ctx context.Context, question *models.Question) error {
	fmt.Println("Inserting question...")
	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO questions (id, question, answer, test_id) VALUES ($1, $2, $3, $4)",
		question.Id, question.Question, question.Answer, question.TestId,
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}
