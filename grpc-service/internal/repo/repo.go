package repo

import (
	"context"

	"github.com/artromone/lccs/grpc-service/internal/models"
)

type BookRepository interface {
	Create(ctx context.Context, book *models.Book) error
}
