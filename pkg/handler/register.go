package handler

import (
	"net/http"
	"text/template"
)

func Register(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/register.html"))
	t.Execute(w, nil)
}