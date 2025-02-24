package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	"github.com/artromone/lccs/grpc-service/internal/repo/postgres"
	"github.com/artromone/lccs/grpc-service/internal/service"
	"github.com/artromone/lccs/grpc-service/internal/transport/grpctransport"
	"github.com/artromone/lccs/grpc-service/proto"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func applyMigrations(connStr string) {
	m, err := migrate.New(
		"file://migrations",
		connStr,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connStr := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	applyMigrations(connStr)

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
