package handler

import (
	"forum/pkg/datamanagement"
	"net/http"
	"text/template"
)

type AccountPage struct {
	Username                  string
	Email                     string
	Profile_picture           string
	Description               string
	FirstName                 string
	LastName                  string
	IsNotValidchangedPwd      bool
	IsNotValidchangedBIO      bool
	IsNotValidEditMail        bool
	IsNotValidchangedUsername bool
}

func setDisplayStructAccount(displayStructAccountPage AccountPage, currentUser datamanagement.Users) AccountPage {
	displayStructAccountPage.Username = currentUser.Username
	displayStructAccountPage.Email = currentUser.Email
	displayStructAccountPage.Profile_picture = currentUser.ProfilePicture
	displayStructAccountPage.Description = currentUser.Description
	displayStructAccountPage.FirstName = currentUser.Firstname
	displayStructAccountPage.LastName = currentUser.Lastname
	return displayStructAccountPage
}

func setDefaultValue(displayStructAccountPage AccountPage) AccountPage {
	displayStructAccountPage.Profile_picture = "../img/PP_wb.png"
	displayStructAccountPage.Username = "Guest"
	displayStructAccountPage.Email = "No email for guest"
	displayStructAccountPage.Description = "No bio added"
	displayStructAccountPage.FirstName = "No First Name"
	displayStructAccountPage.LastName = "No Last Name"
	return displayStructAccountPage
}

func Account(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/account.html", "./static/html/navBar.html"))
	delAccount := r.FormValue("delAccount")
	editMail := r.FormValue("editMail")
	changedPwd1 := r.FormValue("changedPwd1")
	changedPwd2 := r.FormValue("changedPwd2")
	changedBIO := r.FormValue("changedBIO")
	changedFirstname := r.FormValue("changedFirstname")
	changedLastname := r.FormValue("changedLastname")
	changedUsername := r.FormValue("changedUsername")
	cookieIdUser, _ := r.Cookie("idUser")
	cookieIsConnected, _ := r.Cookie("isConnected")
	displayStructAccountPage := AccountPage{}
	idUser := getCookieValue(cookieIdUser)
	isConnected := getCookieValue(cookieIsConnected)
	switch true {
	case delAccount != "":
		datamanagement.AddDeleteUpdateDB("DELETE FROM Users WHERE UserID = ?;", idUser)
		break
	case editMail != "":
		if !datamanagement.IsUserExist(editMail, "") {
			datamanagement.AddDeleteUpdateDB("UPDATE Users SET Email = ? WHERE UserID = ?;", editMail, idUser)
		} else {
			displayStructAccountPage.IsNotValidEditMail = true
		}
		break
	case changedPwd1 != "" && changedPwd2 != "":
		if changedPwd1 == changedPwd2 {
			datamanagement.AddDeleteUpdateDB("UPDATE Users SET Password = ? WHERE UserID = ?;", changedPwd1, idUser)
		} else {
			displayStructAccountPage.IsNotValidchangedPwd = true
		}
		break
	case changedBIO != "":
		datamanagement.AddDeleteUpdateDB("UPDATE Users SET Description = ? WHERE UserID = ?;", changedBIO, idUser)
		break
	case changedFirstname != "":
		datamanagement.AddDeleteUpdateDB("UPDATE Users SET Firstname = ? WHERE UserID = ?;", changedFirstname, idUser)
		break
	case changedLastname != "":
		datamanagement.AddDeleteUpdateDB("UPDATE Users SET Lastname = ? WHERE UserID = ?;", changedLastname, idUser)
		break
	case changedUsername != "":
		if !datamanagement.IsUserExist("", changedUsername) {
			datamanagement.AddDeleteUpdateDB("UPDATE Users SET Username = ? WHERE UserID = ?;", changedUsername, idUser)
		} else {
			displayStructAccountPage.IsNotValidchangedUsername = true
		}
		break
	}
	currentUser := datamanagement.GetUserById(idUser)
	displayStructAccountPage = setDisplayStructAccount(displayStructAccountPage, currentUser)
	displayStructAccountPage.Profile_picture = currentUser.ProfilePicture
	if isConnected != "true" {
		displayStructAccountPage = setDefaultValue(displayStructAccountPage)
	}
	t.ExecuteTemplate(w, "account", displayStructAccountPage)
}
