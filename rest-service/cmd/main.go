package main

import (
	"log"
	"net/http"

	"github.com/artromone/lccs/rest-service/internal/client"
	"github.com/artromone/lccs/rest-service/internal/transport/rest"
)

func main() {
	grpcClient, err := client.NewGRPCClient("grpc-server:50051")
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer grpcClient.Close()

	bookHandler := rest.NewBookHandler(grpcClient)

	http.HandleFunc("/books", bookHandler.CreateBook)
	http.HandleFunc("/books/", bookHandler.GetBook)

	log.Println("REST server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}
}
