package postgres

import (
	"context"
	"database/sql"

	"github.com/artromone/lccs/grpc-service/internal/models"
	"github.com/artromone/lccs/grpc-service/internal/repo"
)

type psqlBookRepositoryImpl struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) repo.BookRepository {
	return &psqlBookRepositoryImpl{db: db}
}

func (r *psqlBookRepositoryImpl) Create(ctx context.Context, book *models.Book) error {
	query := `INSERT INTO books(title, author, status) VALUES($1, $2, $3) RETURNING id`
	return r.db.QueryRowContext(ctx, query, book.Title, book.Author, book.Status).Scan(&book.ID)
}
