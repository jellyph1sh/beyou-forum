package handler

import (
	"crypto/sha256"
	"fmt"
	"forum/pkg/datamanagement"
	"net/http"
	"os/exec"
	"text/template"
	"time"
)

type register struct {
	isValid bool
}

func CreateUser(UserId string, userName string, userFirstName string, userLastName string, userEmail string, stringPasswordInSha256 string) datamanagement.Users {
	nUser := datamanagement.Users{}
	nUser.UserID = UserId
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

func deleteLastByte(tab []byte) []byte {
	returnedTab := []byte{}
	for index, element := range tab {
		if index != len(tab)-1 {
			returnedTab = append(returnedTab, element)
		}
	}
	return returnedTab
}

func Register(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/register.html", "./static/html/navBar.html"))
	userEmail := r.FormValue("email")
	userName := r.FormValue("username")
	userPassword := r.FormValue("password")
	registerDisplay := register{}
	registerDisplay.isValid = true
	if userEmail != "" && userName != "" && userPassword != "" {
		if !datamanagement.IsUserExist(userEmail, userName) {
			newUUID, err := exec.Command("uuidgen").Output()
			if err != nil {
				fmt.Println("user creation error ", err)
			}
			passwordByte := []byte(userPassword)
			passwordInSha256 := sha256.Sum256(passwordByte)
			stringPasswordInSha256 := fmt.Sprintf("%x", passwordInSha256[:])
			fmt.Print(newUUID)
			newUUID = deleteLastByte(newUUID)
			fmt.Print(newUUID)
			nUser := CreateUser(string(newUUID), userName, userName, userName, userEmail, stringPasswordInSha256)
			nDataContainer := datamanagement.DataContainer{}
			nDataContainer.Users = nUser
			datamanagement.AddLineIntoTargetTable(nDataContainer, "Users")
		} else {
			registerDisplay.isValid = false
			fmt.Println("nofvjnorlsfn")
		}
	}

	// fmt.Println(userEmail, userName, userPassword)
	t.ExecuteTemplate(w, "register", nil)
}
