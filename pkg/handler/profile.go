package handler

import (
	"fmt"
	"forum/pkg/datamanagement"
	"net/http"
	"text/template"
)

// pp, pseudo, date de création du compte
// post du mec (date, contenu, nbrLike, nbrDislike)
// topic du mec (ppTopic, nom du topic)
type profile struct {
	UserInfo         datamanagement.Users
	UserCreationDate string
}

func Profile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/profile.html", "./static/html/navBar.html"))
	cookieIdUser, _ := r.Cookie("idUser")
	idUser := getCookieValue(cookieIdUser)
	displayStructProfile := profile{}
	displayStructProfile.UserInfo = datamanagement.GetProfileData(idUser)
	displayStructProfile.UserCreationDate = displayStructProfile.UserInfo.CreationDate.Format("2-1-2006")
	fmt.Println(displayStructProfile.UserInfo.ProfilePicture)
	t.Execute(w, displayStructProfile)
}
