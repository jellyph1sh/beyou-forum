package handler

import (
	"forum/pkg/datamanagement"
	"net/http"
	"text/template"
)

type Moderation struct {
	Reports        []datamanagement.Reports
	Posts          []datamanagement.Posts
	BannedUsers    []datamanagement.Users
	ReportedUsers  []datamanagement.Users
	WordsBlacklist []datamanagement.WordsBlacklist
}

func Automod(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/automod.html", "./static/html/navBar.html"))

	ban := r.FormValue("ban")
	delete := r.FormValue("delete")
	word := r.FormValue("word")
	unban := r.FormValue("unban")
	if word != "" {
		datamanagement.AddWordIntoBlacklist(word)
	} else if unban != "" {
		datamanagement.SetUserStatus(unban, "1")
	} else if ban != "" {
		datamanagement.DeleteReport(ban)
		datamanagement.DeletePost()
	} else if delete != "" {

	}
	t.ExecuteTemplate(w, "automod", Moderation{
		Reports:        datamanagement.GetAllReports(),
		Posts:          datamanagement.GetAllReportedPosts(),
		BannedUsers:    datamanagement.GetAllBannedUsers(),
		ReportedUsers:  datamanagement.GetAllReportedUsers(),
		WordsBlacklist: datamanagement.GetAllBlacklistWords(),
	})
}
