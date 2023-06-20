package handler

import (
	"forum/pkg/datamanagement"
	"net/http"
	"text/template"
)

type invalidPath struct {
	IsConnected string
	IsAdmin     bool
}

func InvalidPath(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/invalidPath.html", "./static/html/navBar.html"))
	cookieConnected, _ := r.Cookie("isConnected")
	IsConnected := getCookieValue(cookieConnected)
	invalidPath := invalidPath{}
	invalidPath.IsConnected = IsConnected
	invalidPath.IsAdmin = false
	if IsConnected == "true" {
		cookieIdUser, _ := r.Cookie("idUser")
		currentUser := datamanagement.GetUserById(getCookieValue(cookieIdUser))
		invalidPath.IsAdmin = currentUser.IsAdmin
	}
	t.ExecuteTemplate(w, "errorPath", IsConnected)
}
