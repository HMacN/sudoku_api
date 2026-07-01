package server

import (
	"net/http"
	"sudoku_api/services/handlers"
)

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /healthy", &handlers.Healthy{})          // TODO: Add handler as separate dependency
	mux.Handle("POST /v1/verify", &handlers.VerifyHandler{}) // TODO: Add handler as separate dependency
	return mux
}

// TODO: For OpenAPI Code Generation: https://www.youtube.com/watch?v=9MuEP01h1XU
