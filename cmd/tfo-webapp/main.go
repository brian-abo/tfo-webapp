package main

import (
	"log"
	"net/http"

	"github.com/brian-abo/tfo-webapp/internal/web"
)

func main() {
	router := web.NewRouter()

	addr := ":8080"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
