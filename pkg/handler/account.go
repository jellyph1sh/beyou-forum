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

// update first name / last name
// update email
// update bio
// update passsword
// update pseudo

// UPDATE Users SET Username = "" WHERE UserID = ""
func Account(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/account.html", "./static/html/navBar.html"))
	delAccount := r.FormValue("delAccount")
	editMail := r.FormValue("editMail")
	changedPwd := r.FormValue("changedPwd")
	cookie, _ := r.Cookie("idUser")
	idUser := getCookieValue(cookie)
	currentUser := datamanagement.GetProfileData(idUser)
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
	if delAccount != "" {
		fmt.Println(delAccount)
	} else if editMail != "" {
		fmt.Println(editMail)
	} else if changedPwd != "" {
		fmt.Println(changedPwd)
	}
	t.ExecuteTemplate(w, "account", p)
}
