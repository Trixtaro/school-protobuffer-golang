package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"trixtaro.dev/go/grpc/database"
	"trixtaro.dev/go/grpc/server"
	"trixtaro.dev/go/grpc/studentpb"
	"trixtaro.dev/go/grpc/testpb"
)

func main() {
	godotenv.Load()

	list, err := net.Listen("tcp", ":5060")

	if err != nil {
		log.Fatal(err)
	}

	connection_url := "postgres://"+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@localhost:5432/"+os.Getenv("DB_NAME")+"?sslmode=disable"
	fmt.Println(connection_url)

	repo, err := database.NewPostgresRepository(connection_url)

	studentServer := server.NewStudentServer(repo)
	testServer := server.NewTestServer(repo)

	fmt.Println("Initializing server...")

	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(grpcServer, studentServer)
	testpb.RegisterTestServiceServer(grpcServer, testServer)

	reflection.Register(grpcServer)

	fmt.Println("Initializing gRPC server...")

	if err := grpcServer.Serve(list); err != nil {
		log.Fatal(err)
	}

	fmt.Println("The end")
}
