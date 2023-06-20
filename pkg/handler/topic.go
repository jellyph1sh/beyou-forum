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
			post.IsLikeByConnectedUser = true
			post.IsDislikeByConnectedUser = false
		} else if datamanagement.IsDislikeByUser(userID, post.PostID) {
			post.IsDislikeByConnectedUser = true
			post.IsLikeByConnectedUser = false
		}
		result = append(result, post)
	}
	return result
}

func isFollowTopic(topicName string, topicDisplayStruct DataTopicPage, idUser string) DataTopicPage {
	if idUser != "" {
		rows := datamanagement.SelectDB("SELECT * FROM Follows WHERE UserID LIKE ? AND TopicID LIKE ?;", idUser, strconv.Itoa(topicDisplayStruct.Topic.TopicID))
		defer rows.Close()
		if !rows.Next() {
			topicDisplayStruct.IsFollow = false
		} else {
			topicDisplayStruct.IsFollow = true
		}
	}
	return topicDisplayStruct
}

func isUpvoteTopic(topicName string, topicDisplayStruct DataTopicPage, idUser string) DataTopicPage {
	if idUser != "" {
		rows := datamanagement.SelectDB("SELECT * FROM Upvotes WHERE UserID LIKE ? AND TopicID LIKE ?;", idUser, strconv.Itoa(topicDisplayStruct.Topic.TopicID))
		defer rows.Close()
		if !rows.Next() {
			topicDisplayStruct.IsUpvote = false
		} else {
			topicDisplayStruct.IsUpvote = true
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
	unLike := r.FormValue("unLike")
	unDislike := r.FormValue("unDislike")
	like := r.FormValue("like")
	dislike := r.FormValue("dislike")
	cookieIsConnected, _ := r.Cookie("isConnected")
	isConnected := getCookieValue(cookieIsConnected)
	topicDisplayStruct := DataTopicPage{}
	topicDisplayStruct.Topic = datamanagement.GetOneTopicByName(topicName)
	topicDisplayStruct = isFollowTopic(topicName, topicDisplayStruct, idUser)
	topicDisplayStruct = isUpvoteTopic(topicName, topicDisplayStruct, idUser)
	if isConnected == "true" {
		switch true {
		case len(newPost) > 0 && len(newPost) <= 500 && datamanagement.CheckContentByBlackListWord(newPost):
			post := datamanagement.DataContainer{Posts: datamanagement.Posts{Content: newPost, AuthorID: idUser, TopicID: topicDisplayStruct.Topic.TopicID, Likes: 0, Dislikes: 0, CreationDate: time.Now(), IsValidPost: true}}
			datamanagement.AddLineIntoTargetTable(post, "Posts")
			http.Redirect(w, r, r.URL.String(), http.StatusSeeOther)
			break
		case clickFollow != "":
			if topicDisplayStruct.IsFollow {
				datamanagement.AddDeleteUpdateDB("DELETE FROM Follows WHERE UserID = ?;", idUser)
				topicDisplayStruct.IsFollow = false
			} else {
				line := datamanagement.DataContainer{Follows: datamanagement.Follows{TopicID: topicDisplayStruct.Topic.TopicID, UserID: idUser}}
				datamanagement.AddLineIntoTargetTable(line, "Follows")
				topicDisplayStruct.IsFollow = true
			}
			break
		case clickUpvote != "":
			if topicDisplayStruct.IsUpvote {
				topicDisplayStruct.IsUpvote = false
				datamanagement.UnUpvotesTopic(topicDisplayStruct.Topic.TopicID, idUser)
			} else {
				topicDisplayStruct.IsUpvote = true
				datamanagement.UpvotesTopic(topicDisplayStruct.Topic.TopicID, idUser)
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
		case unLike != "":
			idPost, _ := strconv.Atoi(unLike)
			datamanagement.UnLikePostManager(idPost, idUser, "unLike")
			break
		case unDislike != "":
			idPost, _ := strconv.Atoi(unDislike)
			datamanagement.UnLikePostManager(idPost, idUser, "unDislike")
			break
		}
	}
	topicDisplayStruct.Posts = transformPostInPostInTopicPage(datamanagement.GetPostByTopic(datamanagement.GetTopicId(topicName)), idUser)
	t := template.Must(template.ParseFiles("./static/html/topic.html"))
	t.Execute(w, topicDisplayStruct)
}
