package web

import (
	"net/http"

	"github.com/brian-abo/tfo-webapp/internal/handler/about"
	"github.com/brian-abo/tfo-webapp/internal/handler/home"
)

// NewRouter creates and configures the HTTP router.
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Static assets
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Home
	mux.HandleFunc("GET /", home.Index)

	// About
	mux.HandleFunc("GET /about", about.Index)

	return mux
}
