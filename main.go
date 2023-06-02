package main

import (
	"fmt"
	"forum/pkg/datamanagement"
	"forum/pkg/handler"
	"net/http"
)

var port = ":8080"

func main() {
	topicTest := datamanagement.DataContainer{Topics: datamanagement.Topics{Title: "topic test", Description: "le topic test", CreatorID: 1, Follows: 1, ValidTopic: true}}
	datamanagement.AddLineIntoTargetTable(topicTest, "Topics")
	//handlers
	http.HandleFunc("/", handler.MainHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("static/fonts"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("static/img"))))
	fmt.Println("(http://localhost"+port+"/home"+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}
