package handler

import (
	"net/http"
	"text/template"
)

func Automod(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/automod.html"))
	t.Execute(w, nil)
}