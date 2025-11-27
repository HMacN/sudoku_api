// swagger:operation POST /v1/solve string Solve
// Returns a placeholder string (for now)
// ---
// responses:
//
//	200: Placeholder
package handlers

import "net/http"

type SolveHandler struct{}

func (s *SolveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Placeholder"))
}
