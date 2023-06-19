package main

import (
	"fmt"
	"forum/pkg/handler"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var port = ":8080"

func main() {
	//handlers
	http.HandleFunc("/", handler.MainHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("static/fonts"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("static/img"))))
	fmt.Println("(http://localhost"+port+"/home"+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}

// datamanagement.AddLineIntoTargetTable(postTest, "Posts")
// nCTN := datamanagement.DataContainer{}
// nPost1 := datamanagement.Topics{}
// nPost1.Title = "test"
// nPost1.Description = ""
// nPost1.Picture = "../img/PP_wb.png"
// nPost1.CreationDate = time.Now()
// nPost1.CreatorID = "c599eb0a-c916-46ae-974d-43593c806438"
// nPost1.Upvotes = 0
// nPost1.Follows = 0
// nPost1.ValidTopic = true
// nCTN.Topics = nPost1
// datamanagement.AddLineIntoTargetTable(nCTN, "Topics")
// nPost1.Title = "test2"
// nPost1.Description = ""
// nPost1.Picture = "../img/PP_wb.png"
// nPost1.CreationDate = time.Now()
// nPost1.CreatorID = "c599eb0a-c916-46ae-974d-43593c806438"
// nPost1.Upvotes = 0
// nPost1.Follows = 0
// nPost1.ValidTopic = true
// nCTN.Topics = nPost1
// datamanagement.AddLineIntoTargetTable(nCTN, "Topics")
// nPost1.Title = "test3"
// nPost1.Description = ""
// nPost1.Picture = "../img/PP_wb.png"
// nPost1.CreationDate = time.Now()
// nPost1.CreatorID = "c599eb0a-c916-46ae-974d-43593c806438"
// nPost1.Upvotes = 0
// nPost1.Follows = 0
// nPost1.ValidTopic = true
// nCTN.Topics = nPost1
// datamanagement.AddLineIntoTargetTable(nCTN, "Topics")
// datamanagement.UpdateLine("UPDATE Topics SET Title='test' WHERE TopicID = 2;")

// nCTN := datamanagement.DataContainer{}
// nPost1 := datamanagement.Posts{}
// nPost1.Content = "zdgfhefsqFDFGJDHSQFSDFR?H"
// nPost1.AuthorID = "c599eb0a-c916-46ae-974d-43593c806438"
// nPost1.TopicID = 1
// nPost1.Likes = 0
// nPost1.Dislikes = 0
// nPost1.CreationDate = time.Now()
// nPost1.IsValidPost = true
// nCTN.Posts = nPost1
// datamanagement.AddLineIntoTargetTable(nCTN, "Posts")
