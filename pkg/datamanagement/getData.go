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
		case table == "User":
			row.Scan(&line.User.ID, &line.User.Name, &line.User.First_name, &line.User.User_name, &line.User.Email, &line.User.Password, &line.User.Is_admin, &line.User.Is_valid, &line.User.Description, &line.User.Profile_image, &line.User.Creation_date)
			break
		case table == "Post":
			var comentary string
			row.Scan(&line.Post.ID, &line.Post.Like, &line.Post.Author_id, &line.Post.Is_valid, &line.Post.Content, comentary, &line.Post.Dislike, &line.Post.Topic, &line.Post.Date)
			if len(comentary) != 0 {
				err := json.Unmarshal([]byte(comentary), &line.Post.Comentary)
				if err != nil {
					fmt.Println(err)
				}
			}
			break
		case table == "Topic":
			var follow string
			row.Scan(&line.Topic.ID, &line.Topic.Title, &line.Topic.Description, &line.Topic.Is_valid, follow, &line.Topic.Creator, &line.Topic.Like)
			if len(follow) != 0 {
				err := json.Unmarshal([]byte(follow), &line.Topic.Follow)
				if err != nil {
					fmt.Println(err)
				}
			}
			break
		case table == "Tag":
		}
	}

	return result
}
