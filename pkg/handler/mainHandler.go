package handler

import (
	"net/http"
	"strings"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.String(), "/")
	if url[1] == "profile" && len(url) > 2 {
		Profile(w, r)
	}
	switch r.URL.Path {
	case "/account":
		Account(w, r)
	case "/automod":
		Automod(w, r)
	case "/explore":
		Explore(w, r)
	case "/home":
		Home(w, r)
	case "/login":
		Login(w, r)
	case "/register":
		Register(w, r)
	default:
		InvalidPath(w, r)
	}
}
