package service

import (
	"context"

	"github.com/artromone/lccs/grpc-service/internal/models"
	"github.com/artromone/lccs/grpc-service/internal/repo"
)

type BookService struct {
	repo repo.BookRepository
}

func NewBookService(repo repo.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(ctx context.Context, book *models.Book) error {
	return s.repo.Create(ctx, book)
}
