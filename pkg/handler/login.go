package handler

import (
	"forum/pkg/datamanagement"
	"net/http"
	"text/template"
	"time"
)

type login struct {
	IsNotValid bool
}

func Login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/login.html", "./static/html/navBar.html"))
	userInput := r.FormValue("userInput")
	userPassword := r.FormValue("userPassword")
	rememberMe := r.FormValue("rememberMe")
	cookieRemeberMe, _ := r.Cookie("Remember")
	CrememberMe := getCookieValue(cookieRemeberMe)
	loginDisplay := login{false}
	if CrememberMe == "true" {
		http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
	}
	if userInput != "" && userPassword != "" {
		ifUserExist, idUser := datamanagement.IsRegister(userInput, userPassword)
		if ifUserExist {
			if rememberMe == "true" {
				cookieIdUser := http.Cookie{Name: "idUser", Value: idUser}
				cookieRemember := http.Cookie{Name: "Remember", Value: "true"}
				http.SetCookie(w, &cookieIdUser)
				http.SetCookie(w, &cookieRemember)
				cookieIsConnected := http.Cookie{Name: "isConnected", Value: "true"}
				http.SetCookie(w, &cookieIsConnected)
			} else {
				cookieIdUser := http.Cookie{Name: "idUser", Value: idUser}
				cookieRememberMe := http.Cookie{Name: "Remember", Value: "false"}
				http.SetCookie(w, &cookieIdUser)
				http.SetCookie(w, &cookieRememberMe)
				cookieIsConnected := http.Cookie{Name: "isConnected", Value: "true", Expires: time.Now().Add(6 * time.Hour)}
				http.SetCookie(w, &cookieIsConnected)
			}
			http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)
		} else {
			loginDisplay.IsNotValid = true
		}
	}
	t.ExecuteTemplate(w, "login", loginDisplay)
}
