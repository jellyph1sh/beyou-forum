package handler

import (
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.String(), "/") // take url
	switch true {
	case url[1] == "account" && len(url) == 2:
		Account(w, r)
	case url[1] == "automod" && len(url) == 2:
		Automod(w, r)
	case url[1] == "explore" && len(url) == 2:
		Explore(w, r)
	case url[1] == "home" && len(url) == 2:
		Home(w, r)
	case url[1] == "login" && len(url) == 2:
		Login(w, r)
	case url[1] == "register" && len(url) == 2:
		Register(w, r)
	case url[1] == "profile" && len(url) == 2:
		Profile(w, r, true)
	case url[1] == "profile" && len(url) > 2:
		Profile(w, r, false)
	default:
		InvalidPath(w, r)
  }
}