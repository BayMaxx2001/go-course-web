package main

import (
	"net/http"

	"github.com/BayMaxx2001/go-course-web/pkg/config"
	"github.com/BayMaxx2001/go-course-web/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// pat
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	//go-chi
	mux := chi.NewRouter()

	//middleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSruf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
