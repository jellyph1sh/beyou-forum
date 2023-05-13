package handler

import (
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
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
	case "/profile":
		Profile(w, r)
	case "/register":
		Register(w, r)
	case "/topic":
		Topic(w, r)
	default:
		InvalidPath(w, r)
	}
}
