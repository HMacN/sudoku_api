package server

import (
	"net/http"
	"sudoku_api/services/handlers"
)

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /healthy", &handlers.Healthy{})  // TODO: Add handler as separate dependency
	mux.Handle("POST /v1/solve", &handlers.Solve{})  // TODO: Add handler as separate dependency
	mux.Handle("POST /v1/solve/", &handlers.Solve{}) // TODO: Add handler as separate dependency
	return mux
}
