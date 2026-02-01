package web

import (
	"database/sql"
	"net/http"

	"github.com/brian-abo/tfo-webapp/internal/handler/about"
	contactHandler "github.com/brian-abo/tfo-webapp/internal/handler/contact"
	"github.com/brian-abo/tfo-webapp/internal/handler/gallery"
	"github.com/brian-abo/tfo-webapp/internal/handler/home"
	"github.com/brian-abo/tfo-webapp/internal/repository"
)

// NewRouter creates and configures the HTTP router.
func NewRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	// Repositories
	contactRepo := repository.NewContactRepository(db)

	// Handlers
	contact := contactHandler.NewHandler(contactRepo)

	// Static assets
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Home
	mux.HandleFunc("GET /", home.Index)

	// About
	mux.HandleFunc("GET /about", about.Index)

	// Contact
	mux.HandleFunc("GET /contact", contact.Index)
	mux.HandleFunc("POST /contact", contact.Submit)

	// Gallery
	mux.HandleFunc("GET /gallery", gallery.Index)

	return mux
}
