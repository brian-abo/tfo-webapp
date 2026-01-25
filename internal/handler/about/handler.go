package about

import (
	"log"
	"net/http"

	"github.com/brian-abo/tfo-webapp/web/features/about"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if err := about.Page().Render(r.Context(), w); err != nil {
		log.Printf("render error: %v", err)
	}
}
