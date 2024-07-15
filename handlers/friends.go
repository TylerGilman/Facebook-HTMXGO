package handlers

import (
	"facebookhtmx/views/friends"
	"log"
	"net/http"
)

func HandleFriends(w http.ResponseWriter, r *http.Request) error {
	r = setHtmxContext(r)
	log.Printf("HX-Request: %s", r.Context().Value(HtmxRequestKey))
	return Render(w, r, friends.Friends())
}
