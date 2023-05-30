package main

import (
	"fmt"
	"forum/pkg/datamanagement"
	"forum/pkg/handler"
	"net/http"
)

var port = ":8080"

func main() {
	// topicTest := datamanagement.DataContainer{Topic: datamanagement.Topic{ID: 1, Title: "topic test", Description: "le topic test", Is_valid: true, Follow: []int{1}, Creator: 1}}
	// postTest := datamanagement.DataContainer{Post: datamanagement.Post{ID: 1, Like: []int{1, 5}, Author_id: 1, Is_valid: true, Content: "Post test", Comentary: []int{}, Dislike: []int{}, Topic: 1}}
	// tagTest := datamanagement.DataContainer{Tag: datamanagement.Tag{ID: 1, Title: "tagTitle", Creator: 1, Like: []int{5, 1}}}
	// userTest := datamanagement.DataContainer{User: datamanagement.User{ID: 1, Name: "admin", First_name: "admin", User_name: "admin", Email: "admin@gmail.com", Password: "pwd", Is_admin: true, Is_valid: true, Description: "the admin", Profile_image: "admin_img", Creation_date: time.Now(), Post_like: []int{}, Post_dislike: []int{}, Topic_like: []int{}}}
	// datamanagement.AddLineIntoTargetTable(userTest, "User", 14)
	// fmt.Println(datamanagement.GetDataForOnePost(1))
	// profileTest := datamanagement.GetProfileData(1)
	// fmt.Println(profileTest.User_name, profileTest.Email, profileTest.Profile_image)
	// sortTopic := datamanagement.GetSortTopic()
	// fmt.Println(sortTopic[0].Like, sortTopic[1].Like, sortTopic[2].Like)
	//handlers
	http.HandleFunc("/", handler.MainHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("static/fonts"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("static/img"))))
	fmt.Println("(http://localhost"+port+"/home"+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}
