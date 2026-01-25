package contact

import (
	"log"
	"net/http"

	"github.com/brian-abo/tfo-webapp/web/features/contact"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if err := contact.Page().Render(r.Context(), w); err != nil {
		log.Printf("render error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Submit handles contact form submissions.
// Currently stubbed - returns success without processing.
func Submit(w http.ResponseWriter, r *http.Request) {
	// TODO: Parse form, validate server-side, send email/store in DB
	w.WriteHeader(http.StatusOK)
}
