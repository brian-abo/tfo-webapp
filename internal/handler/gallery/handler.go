package gallery

import (
	"log"
	"net/http"

	"github.com/brian-abo/tfo-webapp/web/features/gallery"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if err := gallery.Page().Render(r.Context(), w); err != nil {
		log.Printf("render error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
