package handler

import (
	"forum/pkg/datamanagement"
	"net/http"
	"text/template"
)

type AccountPage struct {
	Username        string
	Email           string
	Profile_picture string
	Description     string
	FirstName       string
	LastName        string
}

// update first name / last name
// update email
// update bio
// update passsword
// update pseudo

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
	changedPwd := r.FormValue("changedPwd")
	changedBIO := r.FormValue("changedBIO")
	// changedFirstname := r.FormValue("changedFirstname")
	// changedLastname := r.FormValue("changedLastname")
	// changedUsername := r.FormValue("changedUsername")
	cookie, _ := r.Cookie("idUser")
	idUser := getCookieValue(cookie)
	if idUser == "" && uConnected.IsUserConnected {
		idUser = uConnected.IdUser
	}
	switch true {
	case delAccount != "":
		datamanagement.ExecuterQuery("DELETE FROM Users WHERE UserID ='" + idUser + "';")
		break
	case editMail != "":
		datamanagement.ExecuterQuery("UPDATE Users SET Email = '" + editMail + "' WHERE UserID ='" + idUser + "';")
		break
	case changedPwd != "":
		datamanagement.ExecuterQuery("UPDATE Users SET Password = '" + changedPwd + "' WHERE UserID = '" + idUser + "';")
		break
	case changedBIO != "":
		datamanagement.ExecuterQuery("UPDATE Users SET Description = '" + changedBIO + "' WHERE UserID = '" + idUser + "';")
		break
		// case changedFirstname != "":
		// 	datamanagement.ExecuterQuery("UPDATE Users SET Firstname = '" + changedFirstname + "' WHERE UserID = '" + idUser + "';")
		// 	break
		// case changedLastname != "":
		// 	datamanagement.ExecuterQuery("UPDATE Users SET Lastname = '" + changedLastname + "' WHERE UserID = '" + idUser + "';")
		// 	break
		// case Username != "":
		// 	datamanagement.ExecuterQuery("UPDATE Users SET Username = '" + changedUsername + "' WHERE UserID = '" + idUser + "';")
		// 	break
	}
	currentUser := datamanagement.GetProfileData(idUser)
	displayStructAccountPage := AccountPage{currentUser.Username, currentUser.Email, currentUser.ProfilePicture, currentUser.Description, currentUser.Firstname, currentUser.Lastname}
	displayStructAccountPage.Profile_picture = "../img/PP_wb.png"
	if !uConnected.IsUserConnected {
		displayStructAccountPage = setDefaultValue(displayStructAccountPage)
	}
	t.ExecuteTemplate(w, "account", displayStructAccountPage)
}
