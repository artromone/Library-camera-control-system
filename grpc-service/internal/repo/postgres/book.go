package postgres

import (
	"context"
	"database/sql"
	"fmt"

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

func (r *psqlBookRepositoryImpl) Get(ctx context.Context, id string) (*models.Book, error) {
	query := `SELECT id, title, author, status FROM books WHERE id = $1`
	var book models.Book

	err := r.db.QueryRowContext(ctx, query, id).Scan(&book.ID, &book.Title, &book.Author, &book.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}

	return &book, nil
}
