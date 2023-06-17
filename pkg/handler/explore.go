package handler

import (
	"forum/pkg/datamanagement"
	"net/http"
	"strconv"
	"text/template"
)


func Explore(w http.ResponseWriter, r *http.Request) {
	dataToSend := []datamanagement.Topics{}
	cookieFilter, _ := r.Cookie("filter")
	cookiePaging, _ := r.Cookie("paging")
	if getCookieValue(cookieFilter) == "" && getCookieValue(cookiePaging) == "" {
		cookieFilter = &http.Cookie{Name: "filter", Value: "DESC-Upvote"}
		cookiePaging = &http.Cookie{Name: "paging", Value: "1"}
	}
	topic := r.FormValue("topicForm")
	sort := r.FormValue("sort")
	if topic != "" {
		dataToSend = datamanagement.GetTopicByName(topic)
	} else if sort != "" {
		dataToSend = datamanagement.SortTopics(sort)
	}
	pagingInt, _ := strconv.Atoi(getCookieValue(cookiePaging))
	t := template.Must(template.ParseFiles("./static/html/explore.html"))
	if 2+pagingInt > len(dataToSend)-1 {
		t.Execute(w, dataToSend[pagingInt:len(dataToSend)-1])
	} else {
		t.Execute(w, dataToSend[pagingInt:2+pagingInt])
	}
}
