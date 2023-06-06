package handler

import (
	"forum/pkg/datamanagement"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func Topic(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.String(), "/")
	topicName := url[2]
	row := datamanagement.ReadDB("SELECT * FROM Topics WHERE Title = '" + topicName + "';")
	var topic datamanagement.Topics
	for row.Next() {
		row.Scan(&topic.TopicID, &topic.Title, &topic.Description, &topic.Picture, &topic.CreatorID, &topic.Upvotes, &topic.Follows, &topic.ValidTopic)
		row.Close()
	}
	dataToSend := datamanagement.DataTopicPage{Topic: topic}
	row = datamanagement.ReadDB("SELECT * FROM Posts WHERE TopicID = " + strconv.Itoa(topic.TopicID) + ";")
	for row.Next() {
		var post datamanagement.Posts
		row.Scan(&post.PostID, &post.Content, &post.AuthorID, &post.TopicID, &post.Likes, &post.Dislikes, &post.CreationDate, &post.IsValidPost)
		dataToSend.Posts = append(dataToSend.Posts, post)
		dataToSend.Authors = append(dataToSend.Authors, datamanagement.GetUserById(post.AuthorID))
	}
	newPost := r.FormValue("postContent")
	if len(newPost) > 0 && len(newPost) <= 500 && datamanagement.CheckContentByBlackListWord(newPost) {
		post := datamanagement.DataContainer{Posts: datamanagement.Posts{Content: newPost}}
	}
	t := template.Must(template.ParseFiles("./static/html/topic.html"))
	t.Execute(w, dataToSend)
}
