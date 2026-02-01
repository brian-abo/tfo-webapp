package contact

import (
	"log"
	"net/http"
	"net/mail"
	"strings"

	"github.com/brian-abo/tfo-webapp/internal/repository"
	"github.com/brian-abo/tfo-webapp/web/features/contact"
)

// Handler handles contact page requests.
type Handler struct {
	repo *repository.ContactRepository
}

// NewHandler creates a contact Handler with the given repository.
func NewHandler(repo *repository.ContactRepository) *Handler {
	return &Handler{repo: repo}
}

// Index renders the contact page.
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if err := contact.Page().Render(r.Context(), w); err != nil {
		log.Printf("render error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Submit handles contact form submissions.
func (h *Handler) Submit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	name := strings.TrimSpace(r.FormValue("name"))
	email := strings.TrimSpace(r.FormValue("email"))
	message := strings.TrimSpace(r.FormValue("message"))

	if name == "" || email == "" || message == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	if _, err := mail.ParseAddress(email); err != nil {
		http.Error(w, "Invalid email address", http.StatusBadRequest)
		return
	}

	if _, err := h.repo.Insert(r.Context(), name, email, message); err != nil {
		log.Printf("storing contact submission: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
