package grpctransport

import (
	"context"

	book "github.com/artromone/lccs/grpc-service/internal/models"
	"github.com/artromone/lccs/grpc-service/internal/service"
	"github.com/artromone/lccs/grpc-service/proto"
)

type BookServer struct {
	proto.UnimplementedBookServiceServer
	service *service.BookService
}

func NewBookServer(service *service.BookService) *BookServer {
	return &BookServer{service: service}
}

func (s *BookServer) CreateBook(ctx context.Context, req *proto.Book) (*proto.BookResponse, error) {
	book := &book.Book{
		Title:  req.Title,
		Author: req.Author,
		Status: req.Status,
	}

	if err := s.service.CreateBook(ctx, book); err != nil {
		return nil, err
	}

	return &proto.BookResponse{Id: book.ID}, nil
}

func toProtoBook(b *book.Book) *book.Book {
	return &book.Book{
		ID:     b.ID,
		Title:  b.Title,
		Author: b.Author,
		Status: b.Status,
	}
}

func toDBBook(pbBook *book.Book) *book.Book {
	return &book.Book{
		ID:     pbBook.ID,
		Title:  pbBook.Title,
		Author: pbBook.Author,
		Status: pbBook.Status,
	}
}
