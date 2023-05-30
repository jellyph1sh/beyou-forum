package handler

import (
	"fmt"
	"net/http"
	"text/template"
)

type AccountPage struct {
	Username        string
	Email           string
	Profile_picture string
}

func Account(w http.ResponseWriter, r *http.Request) {
	p := AccountPage{Username: "XxDarkSasukexX", Email: "XxDarkSasukexX@gmail.com", Profile_picture: "../img/PP.png"}
	t := template.Must(template.ParseFiles("./static/html/account.html", "./static/html/navBar.html"))
	delAccount := r.FormValue("delAccount")
	editMail := r.FormValue("editMail")
	changedPwd := r.FormValue("changedPwd")
	if delAccount != "" {
		fmt.Println(delAccount)
	} else if editMail != "" {
		fmt.Println(editMail)
	} else if changedPwd != "" {
		fmt.Println(changedPwd)
	}
	t.ExecuteTemplate(w, "account", p)
}
