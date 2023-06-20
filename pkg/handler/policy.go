package handler

import (
	"net/http"
	"text/template"
)

func Policy(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/policy.html", "./static/html/navBar.html"))
	cookieConnected, _ := r.Cookie("isConnected")
	IsConnected := getCookieValue(cookieConnected)
	t.ExecuteTemplate(w, "policy", IsConnected)
}
