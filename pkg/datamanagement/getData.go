package datamanagement

import (
	"encoding/json"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func GetDataForOnePost(idPost int) DataForOnePost {
	result := DataForOnePost{}
	query := "SELECT like,valid_post,content,commentary,dislike,title,user_name FROM Post LEFT JOIN User ON Post.author = User.id LEFT JOIN Topic ON Post.topic = Topic.id WHERE Post.id = " + strconv.Itoa(idPost) + ";"
	row := readDB(query)
	var like string
	var dislike string
	var commentary string
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&like, &result.Is_valid, &result.Content, &commentary, &dislike, &result.TopicName, &result.AuthorName)
	}
	row.Close()
	var likeInt []int
	var dislikeInt []int
	var commentaryInt []int
	err := json.Unmarshal([]byte(like), &likeInt)
	err = json.Unmarshal([]byte(dislike), &dislikeInt)
	err = json.Unmarshal([]byte(commentary), &commentaryInt)
	if err != nil {
		fmt.Println(err)
	}
	result.NbLike = len(likeInt)
	result.NBDislike = len(dislikeInt)
	result.Comentary = commentaryInt
	return result
}

func GetProfileData(idUser int) User {
	result := User{}
	query := "SELECT user_name,email,profile_img FROM User WHERE id = " + strconv.Itoa(idUser) + ";"
	row := readDB(query)
	for row.Next() {
		row.Scan(&result.User_name, &result.Email, &result.Profile_image)
	}
	row.Close()
	return result
}

func GetSortTopic() []Topic {
	result := []Topic{}
	query := "SELECT * FROM Topic ORDER BY like DESC;"
	row := readDB(query)
	for row.Next() {
		var topic Topic
		var follow string
		row.Scan(&topic.ID, &topic.Title, &topic.Description, &topic.Is_valid, &follow, &topic.Creator, &topic.Like)
		if len(follow) != 0 {
			err := json.Unmarshal([]byte(follow), &topic.Follow)
			if err != nil {
				fmt.Println(err)
			}
		}
		result = append(result, topic)
	}
	row.Close()
	return result
}

func GetSortPost() []Post {
	result := []Post{}
	query := "SELECT * FROM Post ORDER BY like - dislike DESC;"
	row := readDB(query)
	for row.Next() {
		var post Post
		var comentary string
		row.Scan(&post.ID, &post.Like, &post.Author_id, &post.Is_valid, &post.Content, comentary, &post.Dislike, &post.Topic, &post.Date)
		if len(comentary) != 0 {
			err := json.Unmarshal([]byte(comentary), &post.Comentary)
			if err != nil {
				fmt.Println(err)
			}
		}
		result = append(result, post)
	}
	row.Close()
	return result
}

func GetUserByName(search string) []User {
	result := []User{}
	query := "SELECT * FROM User WHERE user_name LIKE '%" + search + "%';"
	row := readDB(query)
	for row.Next() {
		var user User
		var Post_like string
		var Post_dislike string
		var Topic_like string
		row.Scan(&user.ID, &user.Name, &user.First_name, &user.User_name, &user.Email, &user.Password, &user.Is_admin, &user.Is_valid, &user.Description, &user.Profile_image, &user.Creation_date, Post_like, Post_dislike, Topic_like)
		if len(Post_like) != 0 {
			err := json.Unmarshal([]byte(Post_like), &user.Post_like)
			if err != nil {
				fmt.Println(err)
			}
		}
		if len(Post_dislike) != 0 {
			err := json.Unmarshal([]byte(Post_dislike), &user.Post_dislike)
			if err != nil {
				fmt.Println(err)
			}
		}
		if len(Topic_like) != 0 {
			err := json.Unmarshal([]byte(Topic_like), &user.Topic_like)
			if err != nil {
				fmt.Println(err)
			}
		}

		result = append(result, user)
	}
	row.Close()
	return result
}

func GetPostByTopic(topic string) []Post {
	result := []Post{}
	query := "SELECT * FROM Post ORDER BY like - dislike DESC WHERE topic LIKE " + topic + ";"
	row := readDB(query)
	for row.Next() {
		var post Post
		var comentary string
		row.Scan(&post.ID, &post.Like, &post.Author_id, &post.Is_valid, &post.Content, comentary, &post.Dislike, &post.Topic, &post.Date)
		if len(comentary) != 0 {
			err := json.Unmarshal([]byte(comentary), &post.Comentary)
			if err != nil {
				fmt.Println(err)
			}
		}
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
