// swagger:operation GET /healthy Healthy
// Returns 200 OK if the server is healthy.
// ---
// responses:
//
//	200: Healthy
package handlers

import "net/http"

type Healthy struct{}

func (s *Healthy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is Healthy"))
}
