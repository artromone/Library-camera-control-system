package postgres

import (
	"context"
	"database/sql"

	"github.com/artromone/lccs/grpc-service/internal/models"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(ctx context.Context, book *models.Book) error {
	query := `INSERT INTO books(title, author, status) VALUES($1, $2, $3) RETURNING id`
	return r.db.QueryRowContext(ctx, query, book.Title, book.Author, book.Status).Scan(&book.ID)
}
