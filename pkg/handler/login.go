package handler

import (
	"fmt"
	"forum/pkg/datamanagement"
	"net/http"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/login.html", "./static/html/navBar.html"))
	userInput := r.FormValue("userInput")
	userPassword := r.FormValue("userPassword")
	if datamanagement.IsUserExist(userInput, userPassword) {
		fmt.Println("il est register")
		// do login
	} else {
		fmt.Println("pas register")
	}
	t.ExecuteTemplate(w, "login", nil)
}
