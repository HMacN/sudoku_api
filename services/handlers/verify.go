// swagger:operation POST /v1/solve string VerifyHandler
// Returns a placeholder string (for now)
// ---
// responses:
//
//	200: Placeholder
package handlers

import "net/http"

type VerifyHandler struct{}

func (s *VerifyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Placeholder"))
}
