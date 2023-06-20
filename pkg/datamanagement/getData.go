package datamanagement

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func GetPostFromUser(idUser string) []Posts {
	rows := SelectDB("SELECT * FROM Posts WHERE AuthorID = ?;", idUser)
	defer rows.Close()
	result := []Posts{}
	for rows.Next() { // Iterate and fetch the records from result cursor
		var post Posts
		rows.Scan(&post.PostID, &post.Content, &post.AuthorID, &post.TopicID, &post.Likes, &post.Dislikes, &post.CreationDate, &post.IsValidPost)
		result = append(result, post)
	}
	return result
}

func GetPostData(idPost int) Posts {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		fmt.Println(err)
		return Posts{}
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM Posts LEFT JOIN Users ON AuthorID = UserID LEFT JOIN Topics ON Posts.TopicID = Topics.TopicID WHERE PostID = ?;", strconv.Itoa(idPost))

	var post Posts
	if err := row.Scan(&post.PostID, &post.Content, &post.AuthorID, &post.TopicID, &post.Likes, &post.Dislikes, &post.CreationDate, &post.IsValidPost); err != nil {
		fmt.Println(err)
		return Posts{}
	}

	return post
}

func GetSortPost() []Posts {
	rows := SelectDB("SELECT * FROM Posts ORDER BY Likes - Dislikes DESC;")
	defer rows.Close()

	result := []Posts{}
	for rows.Next() {
		var post Posts
		rows.Scan(&post.PostID, &post.Content, &post.AuthorID, &post.TopicID, &post.Likes, &post.Dislikes, &post.CreationDate, &post.IsValidPost)
		result = append(result, post)
	}

	return result
}

func SearchUserByName(search string) []Users {
	rows := SelectDB("SELECT * FROM Users WHERE Username LIKE %?%;", search)
	defer rows.Close()

	result := []Users{}
	for rows.Next() {
		var user Users
		rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Description, &user.CreationDate, &user.ProfilePicture, &user.IsAdmin, &user.ValidUser)
		result = append(result, user)
	}
	return result
}

func GetUserByName(userName string) Users {
	rows := SelectDB("SELECT * FROM Users WHERE Username LIKE ?;", userName)
	defer rows.Close()

	var user Users
	for rows.Next() {
		rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Description, &user.CreationDate, &user.ProfilePicture, &user.IsAdmin, &user.ValidUser)
	}
	return user
}

func GetUserByID(userId string) Users {
	rows := SelectDB("SELECT * FROM Users WHERE UserID LIKE ?;", userId)
	defer rows.Close()

	var user Users
	for rows.Next() {
		rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Description, &user.CreationDate, &user.ProfilePicture, &user.IsAdmin, &user.ValidUser)
	}
	return user
}

func GetTopicId(topicName string) string {
	rows := SelectDB("SELECT TopicID FROM Topics WHERE Title LIKE ?;", topicName)
	defer rows.Close()
	var id string
	for rows.Next() {
		rows.Scan(&id)
	}
	return id
}

func GetPostByTopic(topic string) []Posts {
	rows := SelectDB("SELECT * FROM Posts WHERE TopicID LIKE ?;", topic)
	defer rows.Close()

	result := []Posts{}
	for rows.Next() {
		var post Posts
		rows.Scan(&post.PostID, &post.Content, &post.AuthorID, &post.TopicID, &post.Likes, &post.Dislikes, &post.CreationDate, &post.IsValidPost)
		result = append(result, post)
	}

	return result
}

func GetAllFromTable(table string) []DataContainer {
	rows := SelectDB("SELECT * FROM ?;", table)
	defer rows.Close()

	var result []DataContainer
	for rows.Next() {
		var line DataContainer
		switch true {
		case table == "Users":
			rows.Scan(&line.Users.UserID, &line.Users.Username, &line.Users.Email, &line.Users.Password, &line.Users.Firstname, &line.Users.Lastname, &line.Users.Description, &line.Users.CreationDate, &line.Users.ProfilePicture, &line.Users.IsAdmin, &line.Users.ValidUser)
			break
		case table == "Posts":
			rows.Scan(&line.Posts.PostID, &line.Posts.Content, &line.Posts.AuthorID, &line.Posts.TopicID, &line.Posts.Likes, &line.Posts.Dislikes, &line.Posts.CreationDate, &line.Posts.IsValidPost)
			break
		case table == "Topics":
			rows.Scan(&line.Topics.TopicID, &line.Topics.Title, &line.Topics.Description, &line.Topics.Picture, &line.CreationDate, &line.Topics.CreatorID, &line.Topics.Upvotes, &line.Topics.Follows, &line.Topics.ValidTopic)
			break
		case table == "Tags":
			rows.Scan(&line.Tags.TagID, &line.Tags.Title, &line.Tags.CreatorID)
			break
		case table == "Reports":
			rows.Scan(&line.Reports.ReportID, &line.Reports.PostID, &line.Reports.ReportUserID, &line.Reports.Comment)
			break
		case table == "Dislikes":
			rows.Scan(&line.Dislikes.PostID, &line.Dislikes.UserID)
			break
		case table == "Likes":
			rows.Scan(&line.Dislikes.PostID, &line.Dislikes.UserID)
			break
		case table == "Follows":
			rows.Scan(&line.Follows.FollowID, &line.Follows.TopicID, &line.Follows.UserID)
			break
		case table == "TopicsTags":
			rows.Scan(&line.TopicsTags.TopicID, &line.TopicsTags.TagID)
			break
		case table == "Upvotes":
			rows.Scan(&line.Upvotes.TopicID, &line.Upvotes.UserID)
			break
		case table == "WordsBlacklist":
			rows.Scan(&line.WordsBlacklist.WordID, &line.WordsBlacklist.Word)
			break
		}
		result = append(result, line)
	}
	return result
}

/*
typofsort: 'a-z' - 'z-a' - 'DESC-Upvote' - 'ASC-Upvote' - 'creator'
*/

func SortTopics(typOfSort string) []Topics {
	var query string
	switch typOfSort {
	case "a-z":
		query = "SELECT * FROM Topics ORDER BY Title ASC;"
		break
	case "z-a":
		query = "SELECT * FROM Topics ORDER BY Title DESC;"
		break
	case "DESC-Upvote":
		query = "SELECT * FROM Topics ORDER BY Upvotes DESC;"
		break
	case "DESC-Upvote-Home":
		query = "SELECT * FROM Topics ORDER BY Upvotes DESC LIMIT 3;"
		break
	case "ASC-Upvote":
		query = "SELECT * FROM Topics ORDER BY Upvotes ASC;"
		break
	case "creator":
		query = "SELECT * FROM Topics ORDER BY CreatorID DESC;"
		break
	case "default":
		query = "SELECT * FROM Topics;"
		break
	default:
		fmt.Println("invalid type of sort")
		return nil
	}

	rows := SelectDB(query)
	defer rows.Close()

	var result []Topics
	for rows.Next() {
		var line Topics
		rows.Scan(&line.TopicID, &line.Title, &line.Description, &line.Picture, &line.CreationDate, &line.CreatorID, &line.Upvotes, &line.Follows, &line.ValidTopic)
		result = append(result, line)
	}

	return result
}

/*
condition: 'min upvote'-'max upvote'-'creator'-'max follow'-'min follow'.
refer a number in data for these conditions
*/
func FilterTopics(condition string, data DataFilter) []Topics {
	var query string
	switch condition {
	case "min upvote":
		query = "SELECT * FROM Topics WHERE Upvotes >= ?;"
		break
	case "max upvote":
		query = "SELECT * FROM Topics WHERE Upvotes <= ?;"
		break
	case "creator":
		query = "SELECT * FROM Topics WHERE CreatorID = ?;"
		break
	case "max follow":
		query = "SELECT * FROM Topics WHERE Follows >= ?;"
		break
	case "min follow":
		query = "SELECT * FROM Topics WHERE Follows <= ?;"
		break
	default:
		fmt.Println("Invalid condition")
		return nil
	}

	rows := SelectDB(query, fmt.Sprint(data.Number))
	defer rows.Close()

	var result []Topics
	for rows.Next() {
		var line Topics
		rows.Scan(&line.TopicID, &line.Title, &line.Description, &line.Picture, &line.CreationDate, &line.CreatorID, &line.Upvotes, &line.Follows, &line.ValidTopic)
		result = append(result, line)
	}

	return result
}

func FilterPosts(condition string, data DataFilter) []Posts {
	var query string
	switch condition {
	case "min like":
		query = "SELECT * FROM Posts WHERE Likes >= ?;"
		break
	case "max like":
		query = "SELECT * FROM Posts WHERE Like <= ?;"
		break
	case "min dislike":
		query = "SELECT * FROM Posts WHERE Dislikes >= ?;"
		break
	case "max dislike":
		query = "SELECT * FROM Posts WHERE Dislike <= ?;"
		break
	case "creator":
		query = "SELECT * FROM Posts WHERE CreatorID = ?;"
		break
	case "max follow":
		query = "SELECT * FROM Posts WHERE Follows >= ?;"
		break
	case "min follow":
		query = "SELECT * FROM Posts WHERE Follows <= ?;"
		break
	default:
		fmt.Println("Invalid condition")
		return nil
	}

	rows := SelectDB(query, fmt.Sprint(data.Number))
	defer rows.Close()

	var result []Posts
	for rows.Next() {
		var line Posts
		rows.Scan(&line.PostID, &line.Content, &line.AuthorID, &line.TopicID, &line.Likes, &line.Dislikes, &line.CreationDate, &line.IsValidPost)
		result = append(result, line)
	}

	return result
}

/*
typofsort: 'a-z' - 'z-a' - 'like' - 'dislike' - 'creator'
*/
func SortPosts(typOfSort string) []Posts {
	var query string
	switch typOfSort {
	case "a-z":
		query = "SELECT * FROM Posts ORDER BY Title ASC AND IsValidPost = true;"
		break
	case "z-a":
		query = "SELECT * FROM Posts ORDER BY Title DESC AND IsValidPost = true;"
		break
	case "like":
		query = "SELECT * FROM Posts ORDER BY Likes DESC AND IsValidPost = true;"
		break
	case "dislike":
		query = "SELECT * FROM Posts ORDER BY Dislikes DESC AND IsValidPost = true;"
		break
	case "creator":
		query = "SELECT * FROM Posts ORDER BY CreatorID DESC AND IsValidPost = true;"
		break
	default:
		fmt.Println("invalid type of sort")
		return nil
	}

	rows := SelectDB(query)
	defer rows.Close()

	var result []Posts
	for rows.Next() {
		var line Posts
		rows.Scan(&line.PostID, &line.Content, &line.AuthorID, &line.TopicID, &line.Likes, &line.Dislikes, &line.CreationDate, &line.IsValidPost)
		result = append(result, line)
	}

	return result
}

func GetAllReports() []Reports {
	rows := SelectDB("SELECT * FROM Reports;")
	defer rows.Close()

	var reports []Reports
	for rows.Next() {
		var report Reports
		rows.Scan(&report.ReportID, &report.PostID, &report.ReportUserID, &report.Comment)
		reports = append(reports, report)
	}

	return reports
}

func GetAllReportedUsers() []Users {
	rows := SelectDB("SELECT Users.* FROM Users JOIN Reports ON Users.UserID = Reports.ReportUserID;")
	defer rows.Close()

	var reportedUsers []Users
	for rows.Next() {
		var user Users
		rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Description, &user.CreationDate, &user.ProfilePicture, &user.IsAdmin, &user.ValidUser)
		reportedUsers = append(reportedUsers, user)
	}
	return reportedUsers
}

func GetAllBannedUsers() []Users {
	rows := SelectDB("SELECT * FROM Users WHERE ValidUser = false;")
	defer rows.Close()

	var bannedUsers []Users
	for rows.Next() {
		var user Users
		rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Description, &user.CreationDate, &user.ProfilePicture, &user.IsAdmin, &user.ValidUser)
		bannedUsers = append(bannedUsers, user)
	}

	return bannedUsers
}

func GetAllReportedPosts() []Posts {
	rows := SelectDB("SELECT Posts.* FROM Posts JOIN Reports ON Posts.PostID = Reports.PostID;")
	defer rows.Close()

	var reportedPosts []Posts
	for rows.Next() {
		var post Posts
		rows.Scan(&post.PostID, &post.Content, &post.AuthorID, &post.TopicID, &post.Likes, &post.Dislikes, &post.CreationDate, &post.IsValidPost)
		reportedPosts = append(reportedPosts, post)
	}

	return reportedPosts
}

func GetAllBlacklistWords() []WordsBlacklist {
	rows := SelectDB("SELECT * FROM WordsBlacklist;")
	defer rows.Close()

	var blackListWords []WordsBlacklist
	for rows.Next() {
		var word WordsBlacklist
		rows.Scan(&word.WordID, &word.Word)
		blackListWords = append(blackListWords, word)
	}

	return blackListWords
}

func GetUserById(id string) Users {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		fmt.Println(err)
		return Users{}
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM Users WHERE UserID = ?;", id)

	var user Users
	if err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Description, &user.CreationDate, &user.ProfilePicture, &user.IsAdmin, &user.ValidUser); err != nil {
		fmt.Println(err)
		return Users{}
	}

	return user
}

func GetTopicsById(creatorID string) []Topics {
	rows := SelectDB("SELECT * FROM Topics WHERE CreatorID = ? AND ValidTopic = true;", creatorID)
	defer rows.Close()

	topics := []Topics{}
	for rows.Next() {
		var topic Topics
		rows.Scan(&topic.TopicID, &topic.Title, &topic.Description, &topic.Picture, &topic.CreationDate, &topic.CreatorID, &topic.Upvotes, &topic.Follows, &topic.ValidTopic)
		topics = append(topics, topic)
	}
	return topics
}

func GetTopicsByName(search string) []Topics {
	rows := SelectDB("SELECT * FROM Topics WHERE Title LIKE '%?%';", search)
	defer rows.Close()

	result := []Topics{}
	for rows.Next() {
		var topic Topics
		rows.Scan(&topic.TopicID, &topic.Title, &topic.Description, &topic.Picture, &topic.CreationDate, &topic.CreatorID, &topic.Upvotes, &topic.Follows, &topic.ValidTopic)
		result = append(result, topic)
	}

	return result
}

func GetOneTopicByName(search string) Topics {
	rows := SelectDB("SELECT * FROM Topics WHERE Title = ?;", search)
	defer rows.Close()

	result := Topics{}
	for rows.Next() {
		rows.Scan(&result.TopicID, &result.Title, &result.Description, &result.Picture, &result.CreationDate, &result.CreatorID, &result.Upvotes, &result.Follows, &result.ValidTopic)
	}
	return result
}

func GetTagByName(search string) Tags {
	rows := SelectDB("SELECT * FROM Tags WHERE Title = ?;", search)
	defer rows.Close()

	result := Tags{}
	for rows.Next() {
		rows.Scan(&result.TagID, &result.Title, &result.CreatorID)
	}
	return result
}

func GetTopicByName(topicName string) Topics {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		fmt.Println(err)
		return Topics{}
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM Topics WHERE Title like ?;", topicName)

	var topic Topics
	if err := row.Scan(&topic.TopicID, &topic.Title, &topic.Description, &topic.Picture, &topic.CreationDate, &topic.CreatorID, &topic.Upvotes, &topic.Follows, &topic.ValidTopic); err != nil {
		fmt.Println(err)
		return Topics{}
	}

	return topic
}
