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
	CreationDate string
	CreatorName  string
}

type structDisplayHome struct {
	AllTopics   []TopicsWithUserInfo
	IsConnected string
}

func updateTopicsInTopicsWithUserInfo(topics []datamanagement.Topics) []TopicsWithUserInfo {
	result := []TopicsWithUserInfo{}
	for _, element := range topics {
		// fmt.Println("element.CreatorID", element.CreatorID)
		var topic TopicsWithUserInfo
		topic.TopicID = element.TopicID
		topic.Title = element.Title
		topic.Description = element.Description
		topic.Picture = element.Picture
		topic.CreatorID = element.CreatorID
		topic.Upvotes = element.Upvotes
		topic.Follows = element.Follows
		topic.CreationDate = fmt.Sprint(element.CreationDate.Day()) + " " + fmt.Sprint(element.CreationDate.Month()) + " " + fmt.Sprint(element.CreationDate.Year())
		user := datamanagement.GetUserById(topic.CreatorID)
		topic.CreatorName = user.Username
		// fmt.Println(topic.CreationDate)
		// fmt.Println(element.CreationDate.Day(), element.CreationDate.Month(), element.CreationDate.Year(), element.CreationDate)
		result = append(result, topic)
	}
	return result
}

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/html/home.html", "./static/html/navBar.html", "./static/html/cookiesPopup.html"))
	allTop := datamanagement.SortTopics("DESC-Upvote-Home")
	structDisplayHome := structDisplayHome{}
	structDisplayHome.AllTopics = updateTopicsInTopicsWithUserInfo(allTop)
	cookieConnected, _ := r.Cookie("isConnected")
	IsConnected := getCookieValue(cookieConnected)
	structDisplayHome.IsConnected = IsConnected
	t.ExecuteTemplate(w, "home", structDisplayHome)
}
