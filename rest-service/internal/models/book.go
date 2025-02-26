package models

type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Status string `json:"status"`
}

type BookResponse struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Status string `json:"status"`
}
