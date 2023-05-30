package handler

import (
	"net/http"
	"text/template"
)

func Topic(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/topic.html"))
	t.Execute(w, nil)
}
