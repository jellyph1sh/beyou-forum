package handler

import (
	"forum/pkg/datamanagement"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type DataTopicPage struct {
	Topic    datamanagement.Topics
	Posts    []PostInTopicPage
	IsFollow bool
	IsUpvote bool
}

type PostInTopicPage struct {
	PostID                   int
	Content                  string
	Likes                    int
	Dislikes                 int
	StructuredDate           string
	IsValidPost              bool
	IsLikeByConnectedUser    bool
	IsDislikeByConnectedUser bool
	ProfilePicture           string
	AuthorName               string
}

func transformPostInPostInTopicPage(posts []datamanagement.Posts, userID string) []PostInTopicPage {
	result := []PostInTopicPage{}
	for _, element := range posts {
		var post PostInTopicPage
		post.Content = element.Content
		post.PostID = element.PostID
		post.Likes = element.Likes
		post.Dislikes = element.Dislikes
		post.IsValidPost = element.IsValidPost
		post.StructuredDate = datamanagement.TransformDateInPostFormat(element.CreationDate)
		user := datamanagement.GetUserById(element.AuthorID)
		post.ProfilePicture = user.ProfilePicture
		post.AuthorName = user.Username
		if datamanagement.IsLikeByUser(userID, post.PostID) {
			post.IsDislikeByConnectedUser = true
		} else if datamanagement.IsDislikeByUser(userID, post.PostID) {
			post.IsDislikeByConnectedUser = true
		}
		result = append(result, post)
	}
	return result
}

func isFollowTopic(topicName string, topicDisplayStruct DataTopicPage, idUser string) DataTopicPage {
	if idUser != "" {
		row := datamanagement.ReadDB("SELECT * FROM Follows WHERE UserID LIKE '" + idUser + "' AND TopicID LIKE '" + strconv.Itoa(topicDisplayStruct.Topic.TopicID) + "';")
		row.Close()
		if !row.Next() {
			topicDisplayStruct.IsFollow = false
		} else {
			topicDisplayStruct.IsFollow = true
		}
		row = datamanagement.ReadDB("SELECT * FROM Upvotes WHERE UserID LIKE '" + idUser + "' AND TopicID LIKE '" + strconv.Itoa(topicDisplayStruct.Topic.TopicID) + "';")
		row.Close()
		if !row.Next() {
			topicDisplayStruct.IsFollow = false
		} else {
			topicDisplayStruct.IsFollow = true
		}
	}
	return topicDisplayStruct
}

func Topic(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.String(), "/")
	topicName := url[2]
	cookie, _ := r.Cookie("idUser")
	idUser := getCookieValue(cookie)
	newPost := r.FormValue("postContent")
	clickFollow := r.FormValue("follow")
	clickUpvote := r.FormValue("upvote")
	like := r.FormValue("like")
	dislike := r.FormValue("dislike")
	topicDisplayStruct := DataTopicPage{}
	topicDisplayStruct.Posts = transformPostInPostInTopicPage(datamanagement.GetPostByTopic(datamanagement.GetTopicId(topicName)), idUser)
	topicDisplayStruct.Topic = datamanagement.GetTopicByName(topicName)
	topicDisplayStruct = isFollowTopic(topicName, topicDisplayStruct, idUser)
	cookieIsConnected, _ := r.Cookie("isConnected")
	isConnected := getCookieValue(cookieIsConnected)
	if isConnected == "true" {
		switch true {
		case len(newPost) > 0 && len(newPost) <= 500 && datamanagement.CheckContentByBlackListWord(newPost):
			post := datamanagement.DataContainer{Posts: datamanagement.Posts{Content: newPost, AuthorID: idUser, TopicID: topicDisplayStruct.Topic.TopicID, Likes: 0, Dislikes: 0, CreationDate: time.Now(), IsValidPost: true}}
			datamanagement.AddLineIntoTargetTable(post, "Posts")
			break
		case clickFollow != "":
			if topicDisplayStruct.IsFollow {
				datamanagement.DeleteLineIntoTargetTable("Follows", "UserID = "+idUser)
				topicDisplayStruct.IsFollow = false
			} else {
				line := datamanagement.DataContainer{Follows: datamanagement.Follows{TopicID: topicDisplayStruct.Topic.TopicID, UserID: idUser}}
				datamanagement.AddLineIntoTargetTable(line, "Follows")
				topicDisplayStruct.IsFollow = true
			}
			break
		case clickUpvote != "":
			datamanagement.UpdateUpvotes(topicDisplayStruct.Topic.TopicID, idUser)
			if topicDisplayStruct.IsUpvote {
				topicDisplayStruct.IsUpvote = false
			} else {
				topicDisplayStruct.IsUpvote = true
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
	}
	t := template.Must(template.ParseFiles("./static/html/topic.html"))
	t.Execute(w, topicDisplayStruct)
}
