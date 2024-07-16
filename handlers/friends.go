package handlers

import (
	"facebookhtmx/views/friends"
	"net/http"
)

func HandleFriends(w http.ResponseWriter, r *http.Request) error {
	isHtmxRequest := r.Header.Get("HX-Request") == "true"
	if isHtmxRequest {
		return Render(w, r, friends.Partial())
	} else {
		return Render(w, r, friends.Friends())
	}
}
