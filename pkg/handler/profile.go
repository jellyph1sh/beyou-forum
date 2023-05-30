package handler

import (
	"net/http"
	"text/template"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/profile.html"))
	t.Execute(w, nil)
}