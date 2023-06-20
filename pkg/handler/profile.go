package handler

import (
	"forum/pkg/datamanagement"
	"net/http"
	"strings"
	"text/template"
	"time"
)

type PostWithStructuredDate struct {
	PostID         int
	Content        string
	AuthorID       string
	TopicID        int
	Likes          int
	Dislikes       int
	CreationDate   time.Time
	StructuredDate string
	IsValidPost    bool
	ProfilePicture string
	AuthorName     string
}

// post du mec (date, contenu, nbrLike, nbrDislike)
// topic du mec (ppTopic, nom du topic)
type profile struct {
	UserInfo         datamanagement.Users
	UserCreationDate string
	Posts            []PostWithStructuredDate
	Topics           []datamanagement.Topics
	IsConnected      string
}

func StructureDate(posts []datamanagement.Posts) []PostWithStructuredDate {
	result := []PostWithStructuredDate{}
	for _, element := range posts {
		var post PostWithStructuredDate
		post.Content = element.Content
		post.AuthorID = element.AuthorID
		post.TopicID = element.TopicID
		post.PostID = element.PostID
		post.Likes = element.Likes
		post.Dislikes = element.Dislikes
		post.CreationDate = element.CreationDate
		post.IsValidPost = element.IsValidPost
		post.StructuredDate = datamanagement.TransformDateInPostFormat(post.CreationDate)
		user := datamanagement.GetUserById(post.AuthorID)
		post.ProfilePicture = user.ProfilePicture
		post.AuthorName = user.Username
		result = append(result, post)
	}
	return result
}

func Profile(w http.ResponseWriter, r *http.Request, isMyProfile bool) {
	t := template.Must(template.ParseFiles("./static/html/profile.html", "./static/html/navBar.html"))
	displayStructProfile := profile{}
	url := strings.Split(r.URL.String(), "/")
	if isMyProfile {
		cookieIdUser, _ := r.Cookie("idUser")
		idUser := getCookieValue(cookieIdUser)
		if idUser != "" {
			displayStructProfile.UserInfo = datamanagement.GetUserById(idUser)
		} else {
			displayStructProfile.UserInfo = datamanagement.GetUserByName("guest")
		}
	} else {
		displayStructProfile.UserInfo = datamanagement.GetUserByName(url[2])
	}
	displayStructProfile.UserCreationDate = displayStructProfile.UserInfo.CreationDate.Format("02-01-2006")
	posts := datamanagement.GetPostFromUser(displayStructProfile.UserInfo.UserID)
	displayStructProfile.Topics = datamanagement.GetTopicsById(displayStructProfile.UserInfo.UserID)
	displayStructProfile.Posts = StructureDate(posts)
	cookieConnected, _ := r.Cookie("isConnected")
	IsConnected := getCookieValue(cookieConnected)
	displayStructProfile.IsConnected = IsConnected
	t.ExecuteTemplate(w, "profile", displayStructProfile)
}
