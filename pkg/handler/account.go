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

// 
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
		// "DELETE FROM Users WHERE UserID ='"+ idUser +"';"
		fmt.Println(delAccount)
	} else if editMail != "" {
		//"" UPDATE Users SET Email = '" + editMail +"' WHERE UserID ='"+ idUser +"';"
		fmt.Println(editMail)
	} else if changedPwd != "" {
		//"UPDATE Users SET Password = '" + changedPwd +"' WHERE UserID = '"+ idUser +"';"
		fmt.Println(changedPwd)
	}
		// "UPDATE Users SET Description = '" + changedBIO +"' WHERE UserID = '"+ idUser +"';"
		// "UPDATE Users SET Firstname = '" + changedFirstname +"' WHERE UserID = '"+ idUser +"';"
		// "UPDATE Users SET Lastname = '" + changedLastname +"' WHERE UserID = '"+ idUser +"';"
		// "UPDATE Users SET Username = '" + changedUsername +"' WHERE UserID = '"+ idUser +"';"


	t.ExecuteTemplate(w, "account", p)
}
