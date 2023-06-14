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

type TopicsWithUserInfo struct {
	TopicID      int
	Title        string
	Description  string
	Picture      string
	CreatorID    string
	Upvotes      int
	Follows      int
	ValidTopic   bool
	CreationDate string
	CreatorName  string
}

type structDisplayHome struct {
	AllTopics []TopicsWithUserInfo
}

func updateTopicsInTopicsWithUserInfo(topics []datamanagement.Topics) []TopicsWithUserInfo {
	result := []TopicsWithUserInfo{}
	for _, element := range topics {
		var topic TopicsWithUserInfo
		topic.TopicID = element.TopicID
		topic.Title = element.Title
		topic.Description = element.Description
		topic.Picture = element.Picture
		topic.CreatorID = element.CreatorID
		topic.Upvotes = element.Upvotes
		topic.Follows = element.Follows
		topic.ValidTopic = element.ValidTopic
		topic.CreationDate = fmt.Sprint(element.CreationDate.Day()) + fmt.Sprint(element.CreationDate.Month()) + fmt.Sprint(element.CreationDate.Year())
		user := datamanagement.GetProfileData(topic.CreatorID)
		topic.CreatorName = user.Username
		result = append(result, topic)
	}
	return result
}

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/home.html", "./static/html/navBar.html", "./static/html/cookiesPopup.html"))
	// cookie, _ := r.Cookie("idUser")
	// idUser := getCookieValue(cookie)
	allTop := datamanagement.SortTopics("DESC-Upvote-Home")
	structDisplayHome := structDisplayHome{}
	structDisplayHome.AllTopics = updateTopicsInTopicsWithUserInfo(allTop)
	t.ExecuteTemplate(w, "home", structDisplayHome)
}
