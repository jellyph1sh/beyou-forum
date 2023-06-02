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
	Description     string
}

func Account(w http.ResponseWriter, r *http.Request) {
	currentUser := datamanagement.GetProfileData(1)
	p := AccountPage{currentUser.Email, currentUser.Username, currentUser.ProfilePicture, currentUser.Description}
	if p.Profile_picture == "" {
		p.Profile_picture = "../img/PP_wb.png"
	}
	if p.Username == "" {
		p.Username = "Guest"
	}
	if p.Email == "" {
		p.Email = "No email for guest"
	}
	if p.Description == "" {
		p.Description = "No bio added"
	}
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
