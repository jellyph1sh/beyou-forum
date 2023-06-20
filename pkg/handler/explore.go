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

type DataExplorePage struct {
	Topics       []datamanagement.Topics
	Users        []string
	Upvotes      []string
	CanPrevious  bool
	CanNext      bool
	InvalidTopic bool
}

// return true if some content of the new topic is forbiden
func createTopic(w http.ResponseWriter, r *http.Request, creatorID string) bool {
	title := r.FormValue("topicTitle")
	description := r.FormValue("description")
	tags := r.FormValue("tags")
	if title != "" && (datamanagement.GetOneTopicByName(title) == datamanagement.Topics{}) {
		if !datamanagement.CheckContentByBlackListWord(title) && !datamanagement.CheckContentByBlackListWord(description) && !datamanagement.CheckContentByBlackListWord(tags) && len(strings.Split(title, " ")) == 1 {
			return true
		}
		topic := datamanagement.DataContainer{Topics: datamanagement.Topics{Title: title, Description: description, Picture: "../img/PP_wb.png", CreationDate: time.Now(), CreatorID: creatorID, Upvotes: 0, Follows: 0}}
		datamanagement.AddLineIntoTargetTable(topic, "Topics")
		datamanagement.AddTagsToTopic(tags, creatorID, datamanagement.GetTagByName(title).TagID)
		http.Redirect(w, r, "http://localhost:8080/topic/"+title, http.StatusSeeOther)
	}
	return false
}

func changeUpvote(upvote int) string {
	if upvote < 1000 {
		return fmt.Sprintf("%v", upvote)
	}
	return fmt.Sprintf("%v", upvote/1000.0) + "k"
}

func Explore(w http.ResponseWriter, r *http.Request) {
	cookieUserID, _ := r.Cookie("idUser")
	userId := getCookieValue(cookieUserID)
	dataToSend := DataExplorePage{}
	dataToSend.InvalidTopic = false
	if userId != "" {
		if createTopic(w, r, userId) {
			dataToSend.InvalidTopic = true
		}
	}
	cookieFilter, _ := r.Cookie("filter")
	cookiePaging, _ := r.Cookie("paging")
	if getCookieValue(cookieFilter) == "" && getCookieValue(cookiePaging) == "" {
		cookieFilter = &http.Cookie{Name: "filter", Value: "DESC-Upvote"}
		cookiePaging = &http.Cookie{Name: "paging", Value: "1"}
	}
	topic := r.FormValue("topicSearch")
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
	dataToSend.CanNext = true
	dataToSend.CanPrevious = true
	dataToSend.InvalidTopic = false
	if next != "" && pagingInt*2 < len(dataToSend.Topics) {
		cookiePaging = &http.Cookie{Name: "paging", Value: strconv.Itoa(pagingInt + 1)}
		pagingInt++
	} else if prev != "" && pagingInt > 1 {
		cookiePaging = &http.Cookie{Name: "paging", Value: strconv.Itoa(pagingInt - 1)}
		pagingInt--
	}
	if pagingInt == 1 {
		dataToSend.CanPrevious = false
	}
	t := template.Must(template.ParseFiles("./static/html/explore.html"))
	if len(dataToSend.Topics) == 0 {
		t.Execute(w, dataToSend)
	} else if pagingInt*2 >= len(dataToSend.Topics) {
		dataToSend.CanNext = false
		dataToSend.Upvotes = append(dataToSend.Upvotes, changeUpvote(dataToSend.Topics[(pagingInt-1)*2].Upvotes))
		dataToSend.Users = []string{datamanagement.GetUserById(dataToSend.Topics[(pagingInt-1)*2].CreatorID).Username}
		if pagingInt*2 == len(dataToSend.Topics) {
			dataToSend.Upvotes = append(dataToSend.Upvotes, changeUpvote(dataToSend.Topics[(pagingInt)*2-1].Upvotes))
			dataToSend.Users = append(dataToSend.Users, datamanagement.GetUserById(dataToSend.Topics[(pagingInt)*2-1].CreatorID).Username)
		}
		dataToSend.Topics = dataToSend.Topics[(pagingInt-1)*2:]
		t.Execute(w, dataToSend)
	} else {
		dataToSend.Upvotes = append(dataToSend.Upvotes, changeUpvote(dataToSend.Topics[(pagingInt-1)*2].Upvotes))
		dataToSend.Upvotes = append(dataToSend.Upvotes, changeUpvote(dataToSend.Topics[(pagingInt)*2-1].Upvotes))
		dataToSend.Users = []string{datamanagement.GetUserById(dataToSend.Topics[(pagingInt-1)*2].CreatorID).Username}
		dataToSend.Users = append(dataToSend.Users, datamanagement.GetUserById(dataToSend.Topics[(pagingInt)*2-1].CreatorID).Username)
		dataToSend.Topics = dataToSend.Topics[(pagingInt-1)*2 : pagingInt*2]
		t.Execute(w, dataToSend)
	}
}
