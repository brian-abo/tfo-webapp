package home

import (
	"net/http"

	"github.com/brian-abo/tfo-webapp/web/features/home"
)

func Index(w http.ResponseWriter, r *http.Request) {
	home.Page().Render(r.Context(), w)
}
