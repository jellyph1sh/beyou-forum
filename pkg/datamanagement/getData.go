package datamanagement

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func GetPostFromUser(idUser string) []Posts {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Posts WHERE AuthorID = ?;", idUser)
	if err != nil {
		log.Fatal(err)
		return nil
	}
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
		log.Fatal(err)
		return Posts{}
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM Posts LEFT JOIN Users ON AuthorID = UserID LEFT JOIN Topics ON Posts.TopicID = Topics.TopicID WHERE PostID = ?;", strconv.Itoa(idPost))
	if err != nil {
		log.Fatal(err)
		return Posts{}
	}

	var post Posts
	if err := row.Scan(&post.PostID, &post.Content, &post.AuthorID, &post.TopicID, &post.Likes, &post.Dislikes, &post.CreationDate, &post.IsValidPost); err != nil {
		log.Fatal(err)
		return Posts{}
	}

	return post
}

func GetSortPost() []Posts {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Posts ORDER BY Likes - Dislikes DESC;")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()

	result := []Posts{}
	for rows.Next() {
		var post Posts
		rows.Scan(&post.PostID, &post.Content, &post.AuthorID, &post.TopicID, &post.Likes, &post.Dislikes, &post.CreationDate, &post.IsValidPost)
		result = append(result, post)
	}

	return result
}

func GetUserByName(search string) []Users {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Users WHERE Username LIKE %?%;", search)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()

	result := []Users{}
	for rows.Next() {
		var user Users
		rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Description, &user.CreationDate, &user.ProfilePicture, &user.IsAdmin, &user.ValidUser)
		result = append(result, user)
	}

	return result
}

func GetPostByTopic(topic string) []Posts {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Posts ORDER BY Likes - Dislikes DESC WHERE TopicID LIKE ?;", topic)
	if err != nil {
		log.Fatal(err)
		return nil
	}
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
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM " + table + ";")
	if err != nil {
		log.Fatal(err)
		return nil
	}
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
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

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

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil
	}
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
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

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

	rows, err := db.Query(query, fmt.Sprint(data.Number))
	if err != nil {
		log.Fatal(err)
		return nil
	}
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
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

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

	rows, err := db.Query(query, fmt.Sprint(data.Number))
	if err != nil {
		log.Fatal(err)
		return nil
	}
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
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

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

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()

	var result []Posts
	for rows.Next() {
		var line Posts
		rows.Scan(&line.PostID, &line.Content, &line.AuthorID, &line.TopicID, &line.Likes, &line.Dislikes, &line.CreationDate, &line.IsValidPost)
		result = append(result, line)
	}

	return result
}

func GetUserById(id string) Users {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return Users{}
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM Users WHERE UserID = ?;", id)
	if err != nil {
		log.Fatal(err)
		return Users{}
	}

	var user Users
	if err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Description, &user.CreationDate, &user.ProfilePicture, &user.IsAdmin, &user.ValidUser); err != nil {
		log.Fatal(err)
		return Users{}
	}

	return user
}

func IsPostDLikeByBYser(PostID int, UserID string, DisOrLike string) bool {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM ? WHERE PostID = ? AND UserID = ?", DisOrLike, strconv.Itoa(PostID), UserID)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer rows.Close()

	for rows.Next() {
		return true
	}
	return false
}

func GetTopicsById(creatorID string) []Topics {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

	topics := []Topics{}
	rows, err := db.Query("SELECT * FROM Topics WHERE CreatorID = ? AND ValidTopic = true;", creatorID)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() { // Iterate and fetch the records from result cursor
		var topic Topics
		rows.Scan(&topic.TopicID, &topic.Title, &topic.Description, &topic.Picture, &topic.CreationDate, &topic.CreatorID, &topic.Upvotes, &topic.Follows, &topic.ValidTopic)
		topics = append(topics, topic)
	}

	return topics
}
func GetTopicsByName(search string) []Topics {
	result := []Topics{}
	query := "SELECT * FROM Topics WHERE Title LIKE '%" + search + "%';"
	row := ReadDB(query)
	for row.Next() {
		var topic Topics
		row.Scan(&topic.TopicID, &topic.Title, &topic.Description, &topic.Picture, &topic.CreationDate, &topic.CreatorID, &topic.Upvotes, &topic.Follows, &topic.ValidTopic)
		result = append(result, topic)
	}
	row.Close()
	return result
}

func GetOneTopicByName(search string) Topics {
	result := Topics{}
	query := "SELECT * FROM Topics WHERE Title='" + search + "';"
	row := ReadDB(query)
	for row.Next() {
		row.Scan(&result.TopicID, &result.Title, &result.Description, &result.Picture, &result.CreationDate, &result.CreatorID, &result.Upvotes, &result.Follows, &result.ValidTopic)
		row.Close()
	}
	return result
}
func GetTagByName(search string) Tags {
	result := Tags{}
	query := "SELECT * FROM Tags WHERE Title ='" + search + "';"
	row := ReadDB(query)
	for row.Next() {
		row.Scan(&result.TagID, &result.Title, &result.CreatorID)
		row.Close()
	}
	return result
}
