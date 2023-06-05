package main

import (
	"fmt"
	"forum/pkg/handler"
	"net/http"
)

var port = ":8080"

func main() {
	// topicTest := datamanagement.DataContainer{Topics: datamanagement.Topics{Title: "topic test", Description: "le topic test", CreatorID: 1, Follows: 1, ValidTopic: true}}
	// datamanagement.AddLineIntoTargetTable(topicTest, "Topics")
	// postTest := datamanagement.DataContainer{Posts: datamanagement.Posts{Content: "content posts test", AuthorID: 1, TopicID: 1, Likes: 5, Dislikes: 1, CreationDate: time.Now(), IsValidPost: true}}
	// datamanagement.AddLineIntoTargetTable(postTest, "Posts")
	// datamanagement.LikePostManager(1, 1, "Dislikes")
	// datamanagement.UpdateUpvotes(1, 1)
	// w := datamanagement.DataContainer{WordsBlacklist: datamanagement.WordsBlacklist{WordID: 1, Word: "pd"}}
	// datamanagement.AddLineIntoTargetTable(w, "WordsBlacklist")
	//handlers
	http.HandleFunc("/", handler.MainHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("static/fonts"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("static/img"))))
	fmt.Println("(http://localhost"+port+"/home"+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}
