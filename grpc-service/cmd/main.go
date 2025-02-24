package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/artromone/lccs/grpc-service/internal/repo/postgres"
	"github.com/artromone/lccs/grpc-service/internal/service"
	"github.com/artromone/lccs/grpc-service/internal/transport/grpctransport"
	"github.com/artromone/lccs/grpc-service/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	bookRepo := postgres.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)

	grpcServer := grpc.NewServer()
	bookServer := grpctransport.NewBookServer(bookService)

	proto.RegisterBookServiceServer(grpcServer, bookServer)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
