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
