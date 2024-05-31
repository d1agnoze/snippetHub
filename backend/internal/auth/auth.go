package auth

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

const (
	key    = "phZQTMC5j[`8w;p"
	MaxAge = 86400 * 30
	IsProd = false
)

func NewAuth() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	githubClientID := os.Getenv("GITHUB_CLIENT_ID")
	githubClientSecret := os.Getenv("GITHUB_CLIENT_SECRET")

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store
	goth.UseProviders(
		github.New(githubClientID, githubClientSecret, "http://localhost:8080/auth/github/callback"),
	)
}
