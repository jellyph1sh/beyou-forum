package handler

import (
	"fmt"
	"forum/pkg/datamanagement"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func createTopic(w http.ResponseWriter, r *http.Request, creatorID string) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	tags := r.FormValue("tags")
	if title != "" && datamanagement.CheckContentByBlackListWord(title) && datamanagement.CheckContentByBlackListWord(description) && datamanagement.CheckContentByBlackListWord(tags) && len(strings.Split(title, " ")) == 1 {
		topic := datamanagement.DataContainer{Topics: datamanagement.Topics{Title: title, Description: description, Picture: "../img/PP_wb.png", CreationDate: time.Now(), CreatorID: creatorID, Upvotes: 0, Follows: 0, ValidTopic: true}}
		datamanagement.AddLineIntoTargetTable(topic, "Topics")
		datamanagement.AddTagsToTopic(tags, creatorID, datamanagement.GetOneTopicByName(title).TagID)
		http.Redirect(w, r, "http://localhost:8080/topic/"+title, http.StatusSeeOther)
	}
}

func Explore(w http.ResponseWriter, r *http.Request) {
	cookieUserID, _ := r.Cookie("idUser")
	userId := getCookieValue(cookieUserID)
	if userId != "" {
		createTopic(w, r, userId)
	}
	dataToSend := datamanagement.DataExplorePage{}
	cookieFilter, _ := r.Cookie("filter")
	cookiePaging, _ := r.Cookie("paging")
	if getCookieValue(cookieFilter) == "" && getCookieValue(cookiePaging) == "" {
		cookieFilter = &http.Cookie{Name: "filter", Value: "DESC-Upvote"}
		cookiePaging = &http.Cookie{Name: "paging", Value: "1"}
	}
	topic := r.FormValue("topicForm")
	sort := r.FormValue("sort")
	if topic != "" {
		dataToSend.Topics = datamanagement.GetTopicsByName(topic)
	} else if sort != "" {
		dataToSend.Topics = datamanagement.SortTopics(sort)
	} else {
		dataToSend.Topics = datamanagement.SortTopics("default")
	}
	prev := r.FormValue("previous")
	next := r.FormValue("next")
	pagingInt, _ := strconv.Atoi(getCookieValue(cookiePaging))
	fmt.Println(pagingInt, next, pagingInt*2, len(dataToSend.Topics))
	if next != "" && pagingInt*2 < len(dataToSend.Topics) {
		cookiePaging = &http.Cookie{Name: "paging", Value: strconv.Itoa(pagingInt + 1)}
	} else if prev != "" && pagingInt > 1 {
		cookiePaging = &http.Cookie{Name: "paging", Value: strconv.Itoa(pagingInt - 1)}
	}
	t := template.Must(template.ParseFiles("./static/html/explore.html"))
	if pagingInt+1 > len(dataToSend.Topics) {
		dataToSend.Users = []string{datamanagement.GetUserById(dataToSend.Topics[(pagingInt-1)*2].CreatorID).Username}
		// dataToSend.Users = append(dataToSend.Users, datamanagement.GetUserById(dataToSend.Topics[len(dataToSend.Topics)-1].CreatorID).Username)
		dataToSend.Topics = dataToSend.Topics[(pagingInt-1)*2 : pagingInt*2]
		t.Execute(w, dataToSend)
	} else {
		fmt.Println("test")
		dataToSend.Users = []string{datamanagement.GetUserById(dataToSend.Topics[(pagingInt-1)*2].CreatorID).Username}
		dataToSend.Users = append(dataToSend.Users, datamanagement.GetUserById(dataToSend.Topics[pagingInt].CreatorID).Username)
		dataToSend.Topics = dataToSend.Topics[(pagingInt-1)*2 : pagingInt*2]
		t.Execute(w, dataToSend)
	}
}
