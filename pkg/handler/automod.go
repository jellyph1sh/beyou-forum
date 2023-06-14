package handler

import (
	"forum/pkg/datamanagement"
	"net/http"
	"text/template"
)

type Moderation struct {
	Reports        []datamanagement.Reports
	Posts          []datamanagement.Posts
	Users          []datamanagement.Users
	WordsBlacklist []datamanagement.WordsBlacklist
}

func Automod(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/automod.html", "./static/html/navBar.html"))
	word := r.FormValue("word")
	if word != "" {
		datamanagement.AddWordIntoBlacklist(word)
	}
	t.ExecuteTemplate(w, "automod", Moderation{
		Reports:        datamanagement.GetAllReports(),
		Posts:          datamanagement.GetAllReportedPosts(),
		Users:          datamanagement.GetAllReportedUsers(),
		WordsBlacklist: datamanagement.GetAllBlacklistWords(),
	})
}
