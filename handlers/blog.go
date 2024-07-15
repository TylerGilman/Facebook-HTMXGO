package handlers

import (
	"facebookhtmx/views/blog"
	"log"
	"net/http"
)

func HandleBlog(w http.ResponseWriter, r *http.Request) error {
	r = setHtmxContext(r)
	log.Printf("HX-Request: %s", r.Context().Value(HtmxRequestKey))
	return Render(w, r, blog.Blog())
}
