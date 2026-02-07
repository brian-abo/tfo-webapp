package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/brian-abo/tfo-webapp/internal/auth"
	"github.com/brian-abo/tfo-webapp/internal/config"
	"github.com/brian-abo/tfo-webapp/internal/database"
	"github.com/brian-abo/tfo-webapp/internal/repository"
	"github.com/brian-abo/tfo-webapp/internal/web"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("loading config: %v", err)
	}

	db, err := database.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("connecting to database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("closing database: %v", err)
		}
	}()

	// Repositories
	userRepo := repository.NewUserRepository(db)

	// Auth
	sessionStore := auth.NewPostgresSessionStore(db)
	authMiddleware := auth.NewMiddleware(sessionStore, userRepo)

	var oauthHandler *auth.OAuthHandler
	if cfg.HasFacebookAuth() {
		oauthHandler = auth.NewOAuthHandler(cfg.Facebook, cfg.BaseURL, sessionStore, userRepo)
		log.Printf("Facebook OAuth enabled")
	} else {
		log.Printf("Facebook OAuth not configured (FACEBOOK_CLIENT_ID and FACEBOOK_CLIENT_SECRET required)")
	}

	router := web.NewRouter(db, authMiddleware, oauthHandler)

	log.Printf("listening on %s", cfg.Addr)
	if err := http.ListenAndServe(cfg.Addr, router); err != nil {
		log.Fatal(err)
	}
}
