package handler

import (
	"net/http"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/login.html", "./static/html/navBar.html"))
	t.ExecuteTemplate(w, "login", nil)
}
