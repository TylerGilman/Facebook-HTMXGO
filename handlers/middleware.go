package handlers

import (
	"log"
	"net/http"
)

func AdminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var adminPassword string

		if r.Method == "GET" {
			adminPassword = r.URL.Query().Get("admin_pass")
		} else if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Failed to parse form", http.StatusBadRequest)
				return
			}
			adminPassword = r.FormValue("admin_pass")
		}

		log.Printf("Request URL: %s", r.URL.String())
		log.Printf("Request Method: %s", r.Method)
		log.Printf("Received admin password: %s", adminPassword)

		if adminPassword != "your_secure_password" {
			log.Printf("Unauthorized access attempt")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		log.Printf("Admin access granted")
		next.ServeHTTP(w, r)
	})
}
