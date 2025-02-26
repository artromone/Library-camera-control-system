package client

import (
	"context"
	"log"

	"github.com/artromone/lccs/grpc-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	conn   *grpc.ClientConn
	client proto.BookServiceClient
}

func NewGRPCClient(address string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GRPCClient{
		conn:   conn,
		client: proto.NewBookServiceClient(conn),
	}, nil
}

func (c *GRPCClient) Close() {
	if err := c.conn.Close(); err != nil {
		log.Printf("Failed to close gRPC connection: %v", err)
	}
}

func (c *GRPCClient) CreateBook(ctx context.Context, book *proto.Book) (*proto.BookResponse, error) {
	return c.client.CreateBook(ctx, book)
}

func (c *GRPCClient) GetBook(ctx context.Context, id *proto.BookID) (*proto.Book, error) {
	return c.client.GetBook(ctx, id)
}
