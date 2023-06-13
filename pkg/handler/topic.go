package handler

import (
	"forum/pkg/datamanagement"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func Topic(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.String(), "/")
	topicName := url[2]
	row := datamanagement.ReadDB("SELECT * FROM Topics WHERE Title = '" + topicName + "';")
	var topic datamanagement.Topics
	for row.Next() {
		row.Scan(&topic.TopicID, &topic.Title, &topic.Description, &topic.Picture, &topic.CreatorID, &topic.Upvotes, &topic.Follows, &topic.ValidTopic)
	}
	row.Close()
	dataToSend := datamanagement.DataTopicPage{Topic: topic}
	row = datamanagement.ReadDB("SELECT * FROM Posts WHERE TopicID = " + strconv.Itoa(topic.TopicID) + ";")
	for row.Next() {
		var post datamanagement.Posts
		row.Scan(&post.PostID, &post.Content, &post.AuthorID, &post.TopicID, &post.Likes, &post.Dislikes, &post.CreationDate, &post.IsValidPost)
		dataToSend.Posts = append(dataToSend.Posts, post)
		dataToSend.Authors = append(dataToSend.Authors, datamanagement.GetUserById(post.AuthorID))
	}
	row.Close()
	// add a post
	newPost := r.FormValue("postContent")
	// TO DO : add condition (is user connected)
	if len(newPost) > 0 && len(newPost) <= 500 && datamanagement.CheckContentByBlackListWord(newPost) {
		cookie, _ := r.Cookie("idUser")
		idUser := getCookieValue(cookie)
		//temp condition waiting the cookie
		if len(idUser) == 0 {
			idUser = "1"
		}
		post := datamanagement.DataContainer{Posts: datamanagement.Posts{Content: newPost, AuthorID: idUser, TopicID: dataToSend.Topic.TopicID, Likes: 0, Dislikes: 0, CreationDate: time.Now(), IsValidPost: true}}
		datamanagement.AddLineIntoTargetTable(post, "Posts")
	}
	t := template.Must(template.ParseFiles("./static/html/topic.html"))
	t.Execute(w, dataToSend)
}
