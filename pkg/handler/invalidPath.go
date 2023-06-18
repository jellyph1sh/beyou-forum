package handler

import (
	"net/http"
	"text/template"
)

func InvalidPath(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/invalidPath.html", "./static/html/navBar.html"))
	cookieConnected, _ := r.Cookie("idUser")
	IsConnected := getCookieValue(cookieConnected)
	t.ExecuteTemplate(w, "errorPath", IsConnected)
}
