package main

import (
	"facebookhtmx/handlers"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()
	router.Handle("/*", public())

	// Redirect "/" to "/home"
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home", http.StatusMovedPermanently)
	})

	// Add a new route for "/home"
	router.Get("/home", handlers.Make(handlers.HandleHome))

	router.Get("/friends", handlers.Make(handlers.HandleFriends))
	router.Get("/games", handlers.Make(handlers.HandleGames))
	router.Get("/blog", handlers.Make(handlers.HandleBlog))
	router.Get("/login", handlers.Make(handlers.HandleLoginIndex))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
