package models

type Book struct {
	ID     string `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Status string `db:"status"`
}
