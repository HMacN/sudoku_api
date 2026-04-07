package http

import (
	"net/http"
	handlers2 "sudoku_api/handlers"
)

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /healthy", &handlers2.Healthy{})  // TODO: Add handler as separate dependency
	mux.Handle("POST /v1/solve", &handlers2.Solve{})  // TODO: Add handler as separate dependency
	mux.Handle("POST /v1/solve/", &handlers2.Solve{}) // TODO: Add handler as separate dependency
	return mux
}
