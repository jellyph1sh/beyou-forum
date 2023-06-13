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
	cookie, _ := r.Cookie("idUser")
	idUser := getCookieValue(cookie)

	idUser = "1"

	dataToSend.IsFollow = false
	dataToSend.IsUpvote = false
	if idUser != "" {
		row = datamanagement.ReadDB("SELECT * FROM Follows WHERE UserID = " + idUser + " AND TopicID = " + strconv.Itoa(topic.TopicID) + ";")
		for row.Next() {
			dataToSend.IsFollow = true
			row.Close()
		}
		row = datamanagement.ReadDB("SELECT * FROM Upvotes WHERE UserID = " + idUser + " AND TopicID = " + strconv.Itoa(topic.TopicID) + ";")
		for row.Next() {
			dataToSend.IsUpvote = true
			row.Close()
		}
	}
	// add a post
	newPost := r.FormValue("postContent")
	clickFollow := r.FormValue("follow")
	clickUpvote := r.FormValue("upvote")
	like := r.FormValue("like")
	dislike := r.FormValue("dislike")
	cookieIsConnected, _ := r.Cookie("isConnected")
	isConnected := getCookieValue(cookieIsConnected)
	isConnected = "true"
	if isConnected == "true" {
		switch true {
		case len(newPost) > 0 && len(newPost) <= 500 && datamanagement.CheckContentByBlackListWord(newPost):
			post := datamanagement.DataContainer{Posts: datamanagement.Posts{Content: newPost, AuthorID: idUser, TopicID: dataToSend.Topic.TopicID, Likes: 0, Dislikes: 0, CreationDate: time.Now(), IsValidPost: true}}
			datamanagement.AddLineIntoTargetTable(post, "Posts")
			break
		case clickFollow != "":
			if dataToSend.IsFollow {
				datamanagement.DeleteLineIntoTargetTable("Follows", "UserID = "+idUser)
			} else {
				line := datamanagement.DataContainer{Follows: datamanagement.Follows{TopicID: topic.TopicID, UserID: idUser}}
				datamanagement.AddLineIntoTargetTable(line, "Follows")
			}
			break
		case clickUpvote != "":
			datamanagement.UpdateUpvotes(topic.TopicID, idUser)
			break
		case like != "":
			idPost, _ := strconv.Atoi(like)
			datamanagement.LikePostManager(idPost, idUser, "Likes")
			break
		case dislike != "":
			idPost, _ := strconv.Atoi(dislike)
			datamanagement.LikePostManager(idPost, idUser, "Dislikes")
			break
		}
	}
	t := template.Must(template.ParseFiles("./static/html/topic.html"))
	t.Execute(w, dataToSend)
}
