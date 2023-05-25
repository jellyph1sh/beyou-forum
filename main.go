package main

import (
	"fmt"
	"forum/pkg/datamanagement"
	"forum/pkg/handler"
	"net/http"
)

var port = ":8080"

func main() {
	tagTest := datamanagement.DataContainer{Tag: datamanagement.Tag{ID: 1, Title: "tagTitle", Creator: 1, Like: []int{5, 1}}}
	// userTest := datamanagement.DataContainer{User: datamanagement.User{ID: 1, Name: "admin", First_name: "admin", User_name: "admin", Email: "admin@gmail.com", Password: "pwd", Is_admin: true, Is_valid: true, Description: "the admin", Profile_image: "admin_img", Creation_date: time.Now()}}
	datamanagement.AddLineIntoTargetTable(tagTest, "Tag", 4)
	//handlers
	http.HandleFunc("/", handler.MainHandler)
	fmt.Println("(http://localhost"+port+"/home"+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}
