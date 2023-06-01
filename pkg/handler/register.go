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

func CreateUser(ID int, userName string, userFirstName string, userMainName string, userEmail string, stringPasswordInSha256 string) datamanagement.User {
	nUser := datamanagement.User{}
	nUser.ID = ID
	nUser.Name = userName
	nUser.First_name = userFirstName
	nUser.User_name = userMainName
	nUser.Email = userEmail
	nUser.Password = stringPasswordInSha256
	nUser.Is_admin = false
	nUser.Is_valid = true
	nUser.Description = ""
	nUser.Profile_image = ""
	nUser.Creation_date = time.Now()
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
		nUser := CreateUser(11, userName, userName, userName, userEmail, stringPasswordInSha256)
		nDataContainer := datamanagement.DataContainer{}
		nDataContainer.User = nUser
		// QUERY := "INSERT INTO  User VALUES (7,'" + userName + "','" + userName + "','" + userName + "','" + userEmail + "','" + stringPasswordInSha256 + "'," + "false, true,'','','" + time.Now().String() + "');"
		datamanagement.AddLineIntoTargetTable(nDataContainer, "User", 11)
	} else {
		registerDisplay.isValid = false
		fmt.Println("nofvjnorlsfn")
	}
	// fmt.Println(userEmail, userName, userPassword)
	t.ExecuteTemplate(w, "register", nil)
}
