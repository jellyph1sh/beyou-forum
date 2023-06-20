package handler

import (
	"fmt"
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
	Topics         []datamanagement.Topics
}

func Automod(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/automod.html", "./static/html/navBar.html"))

	banpost := r.FormValue("banpost")
	deletepost := r.FormValue("deletepost")
	cancelreport := r.FormValue("cancelreport")
	word := r.FormValue("word")
	unban := r.FormValue("unban")
	deleteword := r.FormValue("deleteword")
	if word != "" {
		datamanagement.AddWordIntoBlacklist(word)
	} else if unban != "" {
		datamanagement.SetUserStatus(unban, "1")
	} else if banpost != "" {
		datamanagement.DeleteReportFromPost(banpost)
		datamanagement.DeletePost(banpost)
	} else if deletepost != "" {
		datamanagement.DeletePost(deletepost)
	} else if cancelreport != "" {
		datamanagement.DeleteReportFromPost(cancelreport)
	} else if deleteword != "" {
		res := datamanagement.AddDeleteUpdateDB("DELETE FROM WordsBlacklist WHERE Word = ?", deleteword)
		affected, err := res.RowsAffected()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(affected, "deleted!")
	}

	t.ExecuteTemplate(w, "automod", Moderation{
		Reports:        datamanagement.GetAllReports(),
		Posts:          datamanagement.GetAllReportedPosts(),
		BannedUsers:    datamanagement.GetAllBannedUsers(),
		ReportedUsers:  datamanagement.GetAllReportedUsers(),
		WordsBlacklist: datamanagement.GetAllBlacklistWords(),
		Topics:         datamanagement.GetAllReportedTopics(),
	})
}
