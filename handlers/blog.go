package handlers

import (
	"facebookhtmx/views/blog"
	"log"
	"net/http"
)

func HandleBlog(w http.ResponseWriter, r *http.Request) error {
	r = setHtmxContext(r)
	isHtmxRequest := r.Header.Get("HX-Request") == "true"
	log.Printf("HX-Request: %s", r.Context().Value(HtmxRequestKey))

	if isHtmxRequest {
		return Render(w, r, blog.Partial())
	} else {
		return Render(w, r, blog.Blog())
	}
}

func HandleSearch(w http.ResponseWriter, r *http.Request) error {
	query := r.FormValue("query")
	category := r.FormValue("category")

	largeArticles := searchLargeArticles(query, category)
	smallArticles := searchSmallArticles(query, category)

	component := blog.ArticleList(largeArticles, smallArticles)
	return Render(w, r, component)
}

func searchLargeArticles(query string, category string) []blog.Article {
	_ = query
	_ = category
	// Implement your search logic
	return []blog.Article{
		{Title: "Search Result 1", Author: "Author 1", Date: "Jan 30, 2024", Summary: "Summary 1", ImageUrl: "/path/to/image1.jpg"},
		{Title: "Search Result 2", Author: "Author 2", Date: "Jan 31, 2024", Summary: "Summary 2", ImageUrl: "/path/to/image2.jpg"},
	}
}

func searchSmallArticles(query string, category string) []blog.Article {
	_ = query
	_ = category
	// Implement your search logic
	return []blog.Article{
		{Title: "Small Result 1", Category: "Tech", Date: "Feb 1, 2024"},
		{Title: "Small Result 2", Category: "Travel", Date: "Feb 2, 2024"},
		{Title: "Small Result 3", Category: "Food", Date: "Feb 3, 2024"},
	}
}
