package server

import (
	"net/http"

	"gopkg.in/matryer/respond.v1"
)

// mapRoutes defines all the routes the server has
func (s *Server) mapRoutes() {
	// Ping handlers
	s.handleFunc("/ok", func(w http.ResponseWriter, r *http.Request) { respond.With(w, r, http.StatusOK, ok) }).Methods(GET)
}
