package handlers

import (
	"facebookhtmx/views/blog"
	"log/slog"
	"net/http"
)

func HandleBlog(w http.ResponseWriter, r *http.Request) error {
	r = setHtmxContext(r)
	isHtmxRequest := r.Header.Get("HX-Request") == "true"
	slog.Info("HX-Request", "value", r.Context().Value(HtmxRequestKey))

	mainArticles := blog.AllArticles
	sidebarArticles := blog.GetRandomArticles(7) // Get 7 random articles for sidebar

	if isHtmxRequest {
		return Render(w, r, blog.Partial(mainArticles, sidebarArticles))
	} else {
		return Render(w, r, blog.Blog(mainArticles, sidebarArticles))
	}
}

func HandleSearch(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query().Get("query")
	category := r.URL.Query().Get("category")

	slog.Info("Search parameters", "query", query, "category", category)

	searchResults := blog.SearchArticles(query, category)

	// Check if it's an HTMX request
	if r.Header.Get("HX-Request") == "true" {
		// Render only the MainArticles component for HTMX requests
		component := blog.MainArticles(searchResults)
		return Render(w, r, component)
	} else {
		// Render the full page for non-HTMX requests
		mainArticles := searchResults
		sidebarArticles := blog.GetRandomArticles(7) // Get 7 random articles for sidebar
		return Render(w, r, blog.Blog(mainArticles, sidebarArticles))
	}
}
