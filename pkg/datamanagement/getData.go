package datamanagement

import (
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func GetPostData(idPost int) Posts {
	result := Posts{}
	query := "SELECT * FROM Posts LEFT JOIN Users ON AuthorID = UserID LEFT JOIN Topics ON Posts.TopicID = Topics.TopicID WHERE PostID = " + strconv.Itoa(idPost) + ";"
	row := readDB(query)
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&result.PostID, &result.Content, &result.AuthorID, &result.TopicID, &result.Likes, &result.Dislikes, &result.CreationDate, &result.IsValidPost)
	}
	row.Close()
	return result
}

func GetProfileData(idUser int) Users {
	result := Users{}
	query := "SELECT * FROM Users WHERE UserID = " + strconv.Itoa(idUser) + ";"
	row := readDB(query)
	for row.Next() {
		row.Scan(&result.UserID, &result.Username, &result.Email, &result.Password, &result.Firstname, &result.Lastname, &result.Description, &result.CreationDate, &result.ProfilePicture, &result.IsAdmin, &result.ValidUser)
	}
	row.Close()
	return result
}

func GetSortTopic() []Topics {
	result := []Topics{}
	query := "SELECT * FROM Topics ORDER BY Upvotes DESC;"
	row := readDB(query)
	for row.Next() {
		var topic Topics
		row.Scan(&topic.TopicID, &topic.Title, &topic.Description, &topic.CreatorID, &topic.Upvotes, &topic.Follows, &topic.ValidTopic)
		result = append(result, topic)
	}
	row.Close()
	return result
}

func GetSortPost() []Posts {
	result := []Posts{}
	query := "SELECT * FROM Posts ORDER BY Likes - Dislikes DESC;"
	row := readDB(query)
	for row.Next() {
		var post Posts
		row.Scan(&post.PostID, &post.Content, &post.AuthorID, &post.TopicID, &post.Likes, &post.Dislikes, &post.CreationDate, &post.IsValidPost)
		result = append(result, post)
	}
	row.Close()
	return result
}

func GetUserByName(search string) []Users {
	result := []Users{}
	query := "SELECT * FROM Users WHERE Username LIKE '%" + search + "%';"
	row := readDB(query)
	for row.Next() {
		var user Users
		row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Description, &user.CreationDate, &user.ProfilePicture, &user.IsAdmin, &user.ValidUser)
		result = append(result, user)
	}
	row.Close()
	return result
}

func GetPostByTopic(topic string) []Posts {
	result := []Posts{}
	query := "SELECT * FROM Posts ORDER BY Likes - Dislikes DESC WHERE TopicID LIKE " + topic + ";"
	row := readDB(query)
	for row.Next() {
		var post Posts
		row.Scan(&post.PostID, &post.Content, &post.AuthorID, &post.TopicID, &post.Likes, &post.Dislikes, &post.CreationDate, &post.IsValidPost)
		result = append(result, post)
	}
	row.Close()
	return result
}

func GetAllFromTable(table string) []DataContainer {
	row := readDB("SELECT * FROM " + table + ";")
	var result []DataContainer
	for row.Next() {
		var line DataContainer
		switch true {
		case table == "Users":
			row.Scan(&line.Users.UserID, &line.Users.Username, &line.Users.Email, &line.Users.Password, &line.Users.Firstname, &line.Users.Lastname, &line.Users.Description, &line.Users.CreationDate, &line.Users.ProfilePicture, &line.Users.IsAdmin, &line.Users.ValidUser)
			break
		case table == "Posts":
			row.Scan(&line.Posts.PostID, &line.Posts.Content, &line.Posts.AuthorID, &line.Posts.TopicID, &line.Posts.Likes, &line.Posts.Dislikes, &line.Posts.CreationDate, &line.Posts.IsValidPost)
			break
		case table == "Topics":
			row.Scan(&line.Topics.TopicID, &line.Topics.Title, &line.Topics.Description, &line.Topics.CreatorID, &line.Topics.Upvotes, &line.Topics.Follows, &line.Topics.ValidTopic)
			break
		case table == "Tags":
			row.Scan(&line.Tags.TagID, &line.Tags.Title, &line.Tags.CreatorID)
			break
		case table == "Reports":
			row.Scan(&line.Reports.ReportID, &line.Reports.PostID, &line.Reports.ReportUserID, &line.Reports.Comment)
			break
		}
		result = append(result, line)
	}

	return result
}
