package handler

import (
	"fmt"
	"forum/pkg/datamanagement"
	"net/http"
	"text/template"
)

var uConnected datamanagement.UserConnected

func Login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/login.html", "./static/html/navBar.html"))
	userInput := r.FormValue("userInput")
	userPassword := r.FormValue("userPassword")
	rememberMe := r.FormValue("rememberMe")
	ifUserExist, idUser := datamanagement.IsRegister(userInput, userPassword)
	if ifUserExist {
		// expiration := time.Now().Add(365 * 24 * time.Hour)
		// cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
		if rememberMe == "true" {
			cookieIdUser := http.Cookie{Name: "idUser", Value: idUser}
			http.SetCookie(w, &cookieIdUser)
		} else {
			uConnected.IdUser = idUser
			uConnected.IsUserConnected = true
		}
		http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
	} else {
		fmt.Println("pas register")
	}
	t.ExecuteTemplate(w, "login", nil)
}
