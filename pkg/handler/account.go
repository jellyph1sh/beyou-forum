package handler

import (
	"net/http"
	"text/template"
)

func Account(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/account.html"))
	t.Execute(w, nil)
}