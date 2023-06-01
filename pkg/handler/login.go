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
	ifUserExist, idUser := datamanagement.IsRegister(userInput, userPassword)
	if ifUserExist {
		fmt.Println("il est register")
		// do login
		// expiration := time.Now().Add(365 * 24 * time.Hour)
		// cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
		cookie := http.Cookie{Name: "IdUser", Value: string(idUser)}
		http.SetCookie(w, &cookie)
	} else {
		fmt.Println("pas register")
	}
	t.ExecuteTemplate(w, "login", nil)
}
