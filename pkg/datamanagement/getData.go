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
	return result
}
