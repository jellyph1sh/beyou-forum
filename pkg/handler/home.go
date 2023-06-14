package handler

import (
	"fmt"
	"forum/pkg/datamanagement"
	"net/http"
	"text/template"
)

func getCookieValue(cookie *http.Cookie) string {
	var valueReturned string
	test := false
	value := fmt.Sprint(cookie)
	for _, element := range value {
		if test {
			valueReturned += string(element)
		}
		if element == 61 {
			test = true
		}
	}
	return valueReturned
}

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/home.html", "./static/html/navBar.html", "./static/html/cookiesPopup.html"))
	cookie, _ := r.Cookie("idUser")
	idUser := getCookieValue(cookie)
	allTop := datamanagement.SortTopics("DESC-Upvote-Home")
	data := struct {
		AllTopics []datamanagement.Topics
	}{
		AllTopics: allTop,
	}
	fmt.Println(allTop)
	fmt.Println(idUser)
	t.ExecuteTemplate(w, "home", data)
}
