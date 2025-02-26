package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/artromone/lccs/grpc-service/proto"
	"github.com/artromone/lccs/rest-service/internal/client"
	"github.com/artromone/lccs/rest-service/internal/models"
)

type BookHandler struct {
	grpcClient *client.GRPCClient
}

func NewBookHandler(grpcClient *client.GRPCClient) *BookHandler {
	return &BookHandler{grpcClient: grpcClient}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var req models.BookRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return

	}

	grpcReq := &proto.Book{
		Title:  req.Title,
		Author: req.Author,
		Status: req.Status,
	}

	res, err := h.grpcClient.CreateBook(r.Context(), grpcReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.BookResponse{ID: res.Id})
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	if id == "" {
		http.Error(w, "Book ID is required", http.StatusBadRequest)
		return
	}

	grpcReq := &proto.BookID{Id: id}
	res, err := h.grpcClient.GetBook(r.Context(), grpcReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(models.BookResponse{
		ID:     res.Id,
		Title:  res.Title,
		Author: res.Author,
		Status: res.Status,
	})
}
