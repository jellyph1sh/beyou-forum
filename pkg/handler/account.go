package handler

import (
	"fmt"
	"forum/pkg/datamanagement"
	"net/http"
	"text/template"
)

type AccountPage struct {
	Username        string
	Email           string
	Profile_picture string
}

func Account(w http.ResponseWriter, r *http.Request) {
	currentUser := datamanagement.GetProfileData(1)
	p := AccountPage{Username: currentUser.User_name, Email: currentUser.Email, Profile_picture: currentUser.Profile_image}
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
