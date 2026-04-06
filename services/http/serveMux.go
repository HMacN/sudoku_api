package http

import (
	"net/http"
	"sudoku_api/layers/presentation/handlers"
)

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /healthy", &handlers.Healthy{})  // TODO: Add handler as separate dependency
	mux.Handle("POST /v1/solve", &handlers.Solve{})  // TODO: Add handler as separate dependency
	mux.Handle("POST /v1/solve/", &handlers.Solve{}) // TODO: Add handler as separate dependency
	return mux
}
