package handlers

import (
	"facebookhtmx/views/home"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	isHtmxRequest := r.Header.Get("HX-Request") == "true"

	if isHtmxRequest {
		return Render(w, r, home.Partial())
	} else {
		return Render(w, r, home.Index())
	}
}
