package server

import (
	"errors"
	"strings"
	"time"

	"github.com/ysergeyru/go-task-foxg/config"
	"github.com/ysergeyru/go-task-foxg/logger"

	"github.com/patrickmn/go-cache"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	respond "gopkg.in/matryer/respond.v1"
)

// CORS Configuration
var apiCORS = cors.Options{
	AllowedHeaders:     []string{"Accept", "Content-Type", "Content-Length", "X-CRSF-Token", "Authorization", "Cache-Control", "If-Modified-Since", "Pragma", "X-Total-Count"},
	AllowedMethods:     []string{"POST", "GET", "PUT", "OPTIONS", "DELETE"},
	ExposedHeaders:     []string{"X-Total-Count"},
	AllowCredentials:   true,
	OptionsPassthrough: true,
}

// Server is a server
type Server struct {
	logger logger.Logger
	router *mux.Router
	config *config.Config
	cache  *cache.Cache
}

// New creates new server instance
func New(config *config.Config) *Server {
	s := &Server{
		logger: logger.Get(),
		router: mux.NewRouter(),
		config: config,
	}
	// Init in-memory cache with a default expiration time of 60 minutes, and which
	// purges expired items every 80 minutes
	s.cache = cache.New(60*time.Minute, 80*time.Minute)
	if s.cache == nil {
		err := errors.New("Can't init go-cache")
		logger := logger.Get()
		logger.Error(nil, err)
		return nil
	}
	// Map server routes
	s.mapRoutes()
	// Setup CORS
	if len(config.AllowedOrigins) > 0 {
		apiCORS.AllowedOrigins = strings.Split(config.AllowedOrigins, ",")
	}

	return s
}

// HTTPHandler gets the http.Handler for this Server
func (s *Server) HTTPHandler() http.Handler {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Default Secure Headers
		// See https://scotthelme.co.uk/hardening-your-http-response-headers/
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Xss-Protection", "1; mode=block")
		// Breaks IE11 handler when set - header is removed in api.image.handler.go writeImage and writeRaw
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// If it's OPTIONS just allow and retun HTTP Status OK (for CORS)
		if strings.ToUpper(r.Method) == OPTIONS {
			respond.With(w, r, http.StatusOK, ok)
			return
		}

		s.router.ServeHTTP(w, r)
	})
	// Skip CORS if disabled in config
	if config.Get().CORS {
		return cors.New(apiCORS).Handler(handler)
	}
	return handler
}

// ServeHTTP
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// handleFunc passes the requested path to the correct router handler
func (s *Server) handleFunc(path string, fn http.HandlerFunc) *mux.Route {
	return s.router.HandleFunc("/"+strings.TrimLeft(path, "/"), fn)
}
