package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (c *Config) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(c.LoadSessionManager)
	mux.Get("/", c.HomePage)

	mux.Get("/login", c.LoginPage)
	mux.Post("/login", c.PostLoginPage)
	mux.Get("/logout", c.Logout)
	mux.Get("/register", c.RegisterPage)
	mux.Post("/register", c.PostRegisterPage)
	mux.Get("/activate-account", c.ActivateAccount)

	return mux
}
