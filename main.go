package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/TylerGilman/facebookhtmx/handlers"
	"github.com/TylerGilman/facebookhtmx/views/blog"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Initialize the database
	if err := blog.InitDB(); err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer blog.CloseDB() // Ensure the database connection is closed when the program exits

	// Set up the router
	router := chi.NewMux()

	// Static file handling
	router.Handle("/*", public())
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Redirect "/" to "/home"
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home", http.StatusMovedPermanently)
	})

	adminRouter := chi.NewRouter()
	adminRouter.Use(handlers.AdminAuthMiddleware)
	adminRouter.Get("/blog", handlers.Make(handlers.HandleAdminBlogPost))
	adminRouter.Post("/blog/create", handlers.Make(handlers.HandleCreateBlogPost))
	router.Mount("/admin", adminRouter)

	// Public routes
	router.Get("/home", handlers.Make(handlers.HandleHome))
	router.Get("/friends", handlers.Make(handlers.HandleFriends))
	router.Get("/games", handlers.Make(handlers.HandleGames))
	router.Get("/blog", handlers.Make(handlers.HandleBlog))
	router.Get("/blog/search", handlers.Make(handlers.HandleSearch))
	router.Get("/login", handlers.Make(handlers.HandleLoginIndex))

	// Start the server
	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server starting", "listenAddr", listenAddr)
	if err := http.ListenAndServe(listenAddr, router); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
