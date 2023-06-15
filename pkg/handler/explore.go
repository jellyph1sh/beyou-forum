package handler

import (
	"net/http"
	"text/template"
)

func Explore(w http.ResponseWriter, r *http.Request) {


	t := template.Must(template.ParseFiles("./static/html/explore.html"))
	t.Execute(w, nil)
}
