package handler

import (
	"net/http"
	"text/template"
)

func Register(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/register.html", "./static/html/navBar.html"))
	t.ExecuteTemplate(w, "register", nil)
}
