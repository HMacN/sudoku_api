package handlers

import "net/http"

type SolveHandler struct{}

func (s *SolveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Placeholder"))
}
