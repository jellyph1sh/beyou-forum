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
		datamanagement.ExecuterQuery("DELETE FROM Users WHERE UserID ='" + idUser + "';")
		break
	case editMail != "":
		if !datamanagement.IsEmailAlreadyExist(editMail) {
			datamanagement.ExecuterQuery("UPDATE Users SET Email = '" + editMail + "' WHERE UserID ='" + idUser + "';")
		} else {
			displayStructAccountPage.IsNotValidEditMail = true
		}
		break
	case changedPwd1 != "" && changedPwd2 != "":
		if changedPwd1 == changedPwd2 {
			datamanagement.ExecuterQuery("UPDATE Users SET Password = '" + changedPwd1 + "' WHERE UserID = '" + idUser + "';")
		} else {
			displayStructAccountPage.IsNotValidchangedPwd = true
		}
		break
	case changedBIO != "":
		datamanagement.ExecuterQuery("UPDATE Users SET Description = '" + changedBIO + "' WHERE UserID = '" + idUser + "';")
		break
	case changedFirstname != "":
		datamanagement.ExecuterQuery("UPDATE Users SET Firstname = '" + changedFirstname + "' WHERE UserID = '" + idUser + "';")
		break
	case changedLastname != "":
		datamanagement.ExecuterQuery("UPDATE Users SET Lastname = '" + changedLastname + "' WHERE UserID = '" + idUser + "';")
		break
	case changedUsername != "":
		if !datamanagement.IsUsernameAlreadyExist(changedUsername) {
			datamanagement.ExecuterQuery("UPDATE Users SET Username = '" + changedUsername + "' WHERE UserID = '" + idUser + "';")
		} else {
			displayStructAccountPage.IsNotValidchangedUsername = true
		}
		break
	}
	currentUser := datamanagement.GetProfileData(idUser)
	displayStructAccountPage = setDisplayStructAccount(displayStructAccountPage, currentUser)
	displayStructAccountPage.Profile_picture = currentUser.ProfilePicture
	if isConnected != "true" {
		displayStructAccountPage = setDefaultValue(displayStructAccountPage)
	}
	t.ExecuteTemplate(w, "account", displayStructAccountPage)
}
