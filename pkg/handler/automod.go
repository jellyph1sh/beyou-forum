package handler

import (
	"forum/pkg/datamanagement"
	"net/http"
	"text/template"
)

type Moderation struct {
	Reports []datamanagement.Reports
	Posts   []datamanagement.Posts
	Users   []datamanagement.Users
}

func Automod(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/automod.html", "./static/html/navBar.html"))
	t.ExecuteTemplate(w, "automod", Moderation{
		Reports: datamanagement.GetAllReports(),
		Posts:   datamanagement.GetAllReportedPosts(),
		Users:   datamanagement.GetAllReportedUsers(),
	})
}
