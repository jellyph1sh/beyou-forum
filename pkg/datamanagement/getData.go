package datamanagement

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func GetPostFromUser(idUser string) []Posts {
	result := []Posts{}
	query := "SELECT * FROM Posts WHERE AuthorID = '" + idUser + "';"
	row := ReadDB(query)
	for row.Next() { // Iterate and fetch the records from result cursor
		var post Posts
		row.Scan(&post.PostID, &post.Content, &post.AuthorID, &post.TopicID, &post.Likes, &post.Dislikes, &post.CreationDate, &post.IsValidPost)
		result = append(result, post)
	}
	row.Close()
	return result
}

func GetPostData(idPost int) Posts {
	result := Posts{}
	query := "SELECT * FROM Posts LEFT JOIN Users ON AuthorID = UserID LEFT JOIN Topics ON Posts.TopicID = Topics.TopicID WHERE PostID = " + strconv.Itoa(idPost) + ";"
	row := ReadDB(query)
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&result.PostID, &result.Content, &result.AuthorID, &result.TopicID, &result.Likes, &result.Dislikes, &result.CreationDate, &result.IsValidPost)
	}
	row.Close()
	return result
}

func GetProfileData(idUser string) Users {
	result := Users{}
	query := "SELECT * FROM Users WHERE UserID = '" + idUser + "';"
	// fmt.Println("edfghjklm", idUser)
	row := ReadDB(query)
	for row.Next() {
		row.Scan(&result.UserID, &result.Username, &result.Email, &result.Password, &result.Firstname, &result.Lastname, &result.Description, &result.CreationDate, &result.ProfilePicture, &result.IsAdmin, &result.ValidUser)
	}
	row.Close()
	return result
}

func GetSortPost() []Posts {
	result := []Posts{}
	query := "SELECT * FROM Posts ORDER BY Likes - Dislikes DESC;"
	row := ReadDB(query)
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
	row := ReadDB(query)
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
	row := ReadDB(query)
	for row.Next() {
		var post Posts
		row.Scan(&post.PostID, &post.Content, &post.AuthorID, &post.TopicID, &post.Likes, &post.Dislikes, &post.CreationDate, &post.IsValidPost)
		result = append(result, post)
	}
	row.Close()
	return result
}

func GetAllFromTable(table string) []DataContainer {
	row := ReadDB("SELECT * FROM " + table + ";")
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
			row.Scan(&line.Topics.TopicID, &line.Topics.Title, &line.Topics.Description, &line.Topics.Picture, &line.Topics.CreatorID, &line.Topics.Upvotes, &line.Topics.Follows, &line.Topics.ValidTopic)
			break
		case table == "Tags":
			row.Scan(&line.Tags.TagID, &line.Tags.Title, &line.Tags.CreatorID)
			break
		case table == "Reports":
			row.Scan(&line.Reports.ReportID, &line.Reports.PostID, &line.Reports.ReportUserID, &line.Reports.Comment)
			break
		case table == "Dislikes":
			row.Scan(&line.Dislikes.PostID, &line.Dislikes.UserID)
			break
		case table == "Likes":
			row.Scan(&line.Dislikes.PostID, &line.Dislikes.UserID)
			break
		case table == "Follows":
			row.Scan(&line.Follows.FollowID, &line.Follows.TopicID, &line.Follows.UserID)
			break
		case table == "TopicsTags":
			row.Scan(&line.TopicsTags.TopicID, &line.TopicsTags.TagID)
			break
		case table == "Upvotes":
			row.Scan(&line.Upvotes.TopicID, &line.Upvotes.UserID)
			break
		case table == "WordsBlacklist":
			row.Scan(&line.WordsBlacklist.WordID, &line.WordsBlacklist.Word)
			break
		}
		result = append(result, line)
	}
	row.Close()
	return result
}

/*
typofsort: 'a-z' - 'z-a' - 'DESC-Upvote' - 'ASC-Upvote' - 'creator'
*/

func SortTopics(typOfSort string) []Topics {
	var result []Topics
	var row *sql.Rows
	switch typOfSort {
	case "a-z":
		row = ReadDB("SELECT * FROM Topics ORDER BY Title ASC;")
		break
	case "z-a":
		row = ReadDB("SELECT * FROM Topics ORDER BY Title DESC;")
		break
	case "DESC-Upvote":
		row = ReadDB("SELECT TopicID, Title, Description, Picture, CreationDate, CreatorID, Upvotes, Follows, ValidTopic FROM Topics ORDER BY Upvotes DESC;")
		break
	case "DESC-Upvote-Home":
		row = ReadDB("SELECT TopicID, Title, Description, Picture, CreationDate, CreatorID, Upvotes, Follows, ValidTopic FROM Topics ORDER BY Upvotes DESC LIMIT 3;")
		// row = ReadDB("SELECT TopicID, Title, Description, Picture, CreationDate, CreatorID, Upvotes, Follows, ValidTopic FROM Topics;")
		// row = ReadDB("SELECT * FROM Topics ORDER BY Upvotes DESC LIMIT 3;")
		break
	case "ASC-Upvote":
		row = ReadDB("SELECT * FROM Topics ORDER BY Upvotes ASC;")
		break
	case "creator":
		row = ReadDB("SELECT * FROM Topics ORDER BY CreatorID DESC;")
		break
	default:
		fmt.Println("invalid type of sort")
		return result
	}
	for row.Next() {
		var line Topics
		row.Scan(&line.TopicID, &line.Title, &line.Description, &line.Picture, &line.CreationDate, &line.CreatorID, &line.Upvotes, &line.Follows, &line.ValidTopic)
		result = append(result, line)
	}
	row.Close()
	return result
}

/*
condition: 'min upvote'-'max upvote'-'creator'-'max follow'-'min follow'.
refer a number in data for these conditions
*/
func FilterTopics(condition string, data DataFilter) []Topics {
	var result []Topics
	var row *sql.Rows
	switch condition {
	case "min upvote":
		row = ReadDB("SELECT * FROM Topics WHERE Upvotes >= " + fmt.Sprint(data.number) + ";")
		break
	case "max upvote":
		row = ReadDB("SELECT * FROM Topics WHERE Upvotes <= " + fmt.Sprint(data.number) + ";")
		break
	case "creator":
		row = ReadDB("SELECT * FROM Topics WHERE CreatorID = " + fmt.Sprint(data.number) + ";")
		break
	case "max follow":
		row = ReadDB("SELECT * FROM Topics WHERE Follows >= " + fmt.Sprint(data.number) + ";")
		break
	case "min follow":
		row = ReadDB("SELECT * FROM Topics WHERE Follows <= " + fmt.Sprint(data.number) + ";")
		break
	default:
		fmt.Println("Invalid condition")
		return result
	}
	for row.Next() {
		var line Topics
		row.Scan(&line.TopicID, &line.Title, &line.Description, &line.Picture, &line.CreatorID, &line.Upvotes, &line.Follows, &line.ValidTopic)
		result = append(result, line)
	}
	row.Close()
	return result
}

func FilterPosts(condition string, data DataFilter) []Posts {
	var result []Posts
	var row *sql.Rows
	switch condition {
	case "min like":
		row = ReadDB("SELECT * FROM Posts WHERE Likes >= " + fmt.Sprint(data.number) + ";")
		break
	case "max like":
		row = ReadDB("SELECT * FROM Posts WHERE Like <= " + fmt.Sprint(data.number) + ";")
		break
	case "min dislike":
		row = ReadDB("SELECT * FROM Posts WHERE Dislikes >= " + fmt.Sprint(data.number) + ";")
		break
	case "max dislike":
		row = ReadDB("SELECT * FROM Posts WHERE Dislike <= " + fmt.Sprint(data.number) + ";")
		break
	case "creator":
		row = ReadDB("SELECT * FROM Posts WHERE CreatorID = " + fmt.Sprint(data.number) + ";")
		break
	case "max follow":
		row = ReadDB("SELECT * FROM Posts WHERE Follows >= " + fmt.Sprint(data.number) + ";")
		break
	case "min follow":
		row = ReadDB("SELECT * FROM Posts WHERE Follows <= " + fmt.Sprint(data.number) + ";")
		break
	default:
		fmt.Println("Invalid condition")
		return result
	}
	for row.Next() {
		var line Posts
		row.Scan(&line.PostID, &line.Content, &line.AuthorID, &line.TopicID, &line.Likes, &line.Dislikes, &line.CreationDate, &line.IsValidPost)
		result = append(result, line)
	}
	row.Close()
	return result
}

/*
typofsort: 'a-z' - 'z-a' - 'like' - 'dislike' - 'creator'
*/
func SortPosts(typOfSort string) []Posts {
	var result []Posts
	var row *sql.Rows
	switch typOfSort {
	case "a-z":
		row = ReadDB("SELECT * FROM Posts ORDER BY Title ASC;")
		break
	case "z-a":
		row = ReadDB("SELECT * FROM Posts ORDER BY Title DESC;")
		break
	case "like":
		row = ReadDB("SELECT * FROM Posts ORDER BY Likes DESC;")
		break
	case "dislike":
		row = ReadDB("SELECT * FROM Posts ORDER BY Dislikes DESC;")
		break
	case "creator":
		row = ReadDB("SELECT * FROM Posts ORDER BY CreatorID DESC;")
		break
	default:
		fmt.Println("invalid type of sort")
		return result
	}

	for row.Next() {
		var line Posts
		row.Scan(&line.PostID, &line.Content, &line.AuthorID, &line.TopicID, &line.Likes, &line.Dislikes, &line.CreationDate, &line.IsValidPost)
		result = append(result, line)
	}
	row.Close()
	return result
}

func GetUserById(id string) Users {
	row := ReadDB("SELECT * FROM Users WHERE UserID = " + id + ";")
	for row.Next() {
		var user Users
		row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Description, &user.CreationDate, &user.ProfilePicture, &user.IsAdmin, &user.ValidUser)
		row.Close()
		return user
	}
	return Users{}
}

func IsPostDLikeByBYser(PostID int, UserID string, DisOrLike string) bool {
	row := ReadDB("SELECT * FROM " + DisOrLike + " WHERE PostID = " + strconv.Itoa(PostID) + " AND UserID = " + UserID)
	for row.Next() {
		return true
	}
	row.Close()
	return false
}
func GetTopicsById(idUser string) []Topics {
	result := []Topics{}
	query := "SELECT * FROM Topics WHERE CreatorID = '" + idUser + "';"
	row := ReadDB(query)
	for row.Next() { // Iterate and fetch the records from result cursor
		var topic Topics
		row.Scan(&topic.TopicID, &topic.Title, &topic.Description, &topic.Picture, &topic.CreatorID, &topic.Upvotes, &topic.Follows, &topic.ValidTopic)
		result = append(result, topic)
	}
	row.Close()
	return result
}
