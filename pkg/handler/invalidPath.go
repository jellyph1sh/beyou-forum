package handler

import (
	"net/http"
	"text/template"
)

func InvalidPath(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/invalidPath.html"))
	t.Execute(w, nil)
}
