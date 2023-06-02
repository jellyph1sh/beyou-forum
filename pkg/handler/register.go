package handler

import (
	"crypto/sha256"
	"fmt"
	"forum/pkg/datamanagement"
	"net/http"
	"text/template"
	"time"
)

type register struct {
	isValid bool
}

func CreateUser(userName string, userFirstName string, userLastName string, userEmail string, stringPasswordInSha256 string) datamanagement.Users {
	nUser := datamanagement.Users{}
	nUser.Username = userName
	nUser.Email = userEmail
	nUser.Password = stringPasswordInSha256
	nUser.Firstname = userFirstName
	nUser.Lastname = userLastName
	nUser.Description = ""
	nUser.CreationDate = time.Now()
	nUser.ProfilePicture = ""
	nUser.IsAdmin = false
	nUser.ValidUser = true
	return nUser
}

func Register(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/register.html", "./static/html/navBar.html"))
	userEmail := r.FormValue("email")
	userName := r.FormValue("username")
	userPassword := r.FormValue("password")
	registerDisplay := register{}
	registerDisplay.isValid = true

	if !datamanagement.IsUserExist(userEmail, userName) {
		passwordByte := []byte(userPassword)
		passwordInSha256 := sha256.Sum256(passwordByte)
		stringPasswordInSha256 := fmt.Sprintf("%x", passwordInSha256[:])
		nUser := CreateUser(userName, userName, userName, userEmail, stringPasswordInSha256)
		nDataContainer := datamanagement.DataContainer{}
		nDataContainer.Users = nUser
		// QUERY := "INSERT INTO  User VALUES (7,'" + userName + "','" + userName + "','" + userName + "','" + userEmail + "','" + stringPasswordInSha256 + "'," + "false, true,'','','" + time.Now().String() + "');"
		datamanagement.AddLineIntoTargetTable(nDataContainer, "Users", 11)
	} else {
		registerDisplay.isValid = false
		fmt.Println("nofvjnorlsfn")
	}
	// fmt.Println(userEmail, userName, userPassword)
	t.ExecuteTemplate(w, "register", nil)
}
