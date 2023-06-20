package handler

import (
	"fmt"
	"forum/pkg/datamanagement"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func Topic(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.String(), "/")
	topicName := url[2]
	rows := datamanagement.SelectDB("SELECT * FROM Topics WHERE Title = ?;", topicName)
	defer rows.Close()

	var topic datamanagement.Topics
	for rows.Next() {
		rows.Scan(&topic.TopicID, &topic.Title, &topic.Description, &topic.Picture, &topic.CreatorID, &topic.Upvotes, &topic.Follows, &topic.ValidTopic)
	}
	dataToSend := datamanagement.DataTopicPage{}
	cookie, _ := r.Cookie("idUser")
	idUser := getCookieValue(cookie)
	if idUser != "" {
		rows := datamanagement.SelectDB("SELECT * FROM Follows WHERE UserID = ? AND TopicID = ?;", idUser, strconv.Itoa(topic.TopicID))
		defer rows.Close()
		for rows.Next() {
			dataToSend.IsFollow = true
		}
		rows = datamanagement.SelectDB("SELECT * FROM Upvotes WHERE UserID = ? AND TopicID = ?;", idUser, strconv.Itoa(topic.TopicID))
		defer rows.Close()
		for rows.Next() {
			dataToSend.IsUpvote = true
		}
	}
	dataToSend = datamanagement.DataTopicPage{Topic: topic}
	rows = datamanagement.SelectDB("SELECT * FROM Posts WHERE TopicID = ?;", strconv.Itoa(topic.TopicID))
	defer rows.Close()
	for rows.Next() {
		var post datamanagement.Posts
		rows.Scan(&post.PostID, &post.Content, &post.AuthorID, &post.TopicID, &post.Likes, &post.Dislikes, &post.CreationDate, &post.IsValidPost)
		dataToSend.Posts = append(dataToSend.Posts, post)
		dataToSend.Authors = append(dataToSend.Authors, datamanagement.GetUserById(post.AuthorID))
	}
	// idUser = "1" //delete this line
	dataToSend.IsFollow = false
	dataToSend.IsUpvote = false
	// add a post
	newPost := r.FormValue("postContent")
	clickFollow := r.FormValue("follow")
	clickUpvote := r.FormValue("upvote")
	like := r.FormValue("like")
	dislike := r.FormValue("dislike")
	cookieIsConnected, _ := r.Cookie("isConnected")
	isConnected := getCookieValue(cookieIsConnected)
	// isConnected = "true" //delete this line
	if isConnected == "true" {
		switch true {
		case len(newPost) > 0 && len(newPost) <= 500 && datamanagement.CheckContentByBlackListWord(newPost):
			post := datamanagement.DataContainer{Posts: datamanagement.Posts{Content: newPost, AuthorID: idUser, TopicID: dataToSend.Topic.TopicID, Likes: 0, Dislikes: 0, CreationDate: time.Now(), IsValidPost: true}}
			datamanagement.AddLineIntoTargetTable(post, "Posts")
			break
		case clickFollow != "":
			if dataToSend.IsFollow {
				res := datamanagement.AddDeleteUpdateDB("DELETE FROM Follows WHERE UserID = ?;", idUser)
				affected, err := res.RowsAffected()
				if err != nil {
					log.Fatal(err)
					return
				}
				fmt.Println(affected, "deleted!")
				dataToSend.IsFollow = false
			} else {
				line := datamanagement.DataContainer{Follows: datamanagement.Follows{TopicID: topic.TopicID, UserID: idUser}}
				datamanagement.AddLineIntoTargetTable(line, "Follows")
				dataToSend.IsFollow = true
			}
			break
		case clickUpvote != "":
			datamanagement.UpdateUpvotes(topic.TopicID, idUser)
			if dataToSend.IsUpvote {
				dataToSend.IsUpvote = false
			} else {
				dataToSend.IsUpvote = true
			}
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
		for _, p := range dataToSend.Posts {
			dataToSend.Likes = append(dataToSend.Likes, datamanagement.IsPostDLikeByBYser(p.PostID, idUser, "Likes"))
			dataToSend.Dislikes = append(dataToSend.Likes, datamanagement.IsPostDLikeByBYser(p.PostID, idUser, "Dislikes"))
		}
	}
	t := template.Must(template.ParseFiles("./static/html/topic.html"))
	t.Execute(w, dataToSend)
}
