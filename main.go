package main

import (
	"fmt"
	"forum/pkg/handler"

	"net/http"
)

var port = ":8080"

func main() {
	// topicTest := datamanagement.DataContainer{Topics: datamanagement.Topics{Title: "topicTest3", Description: "topicTest", CreatorID: 1, Upvotes: 0, Follows: 0, ValidTopic: true}}
	// postTest := datamanagement.DataContainer{Posts: datamanagement.Posts{Content: "un autre post test odiajdiuoazudanjudfnajudfnaefn", AuthorID: "1", TopicID: 1, Likes: 15, Dislikes: 0, CreationDate: time.Now(), IsValidPost: true}}
	// userTest := datamanagement.DataContainer{Users: datamanagement.Users{UserID: "1", Username: "DarkSasuke", Email: "email", Password: "oipjfziiofnziofnez", Firstname: "Dark", Lastname: "Link", Description: "baka", CreationDate: time.Now(), ProfilePicture: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQLeR3wX3QrJpZGlb6fIeU-XPPRgxGlP5coYmQaBFZJ&s", IsAdmin: false, ValidUser: true}}
	// datamanagement.AddLineIntoTargetTable(postTest, "Posts")
	// nCTN := datamanagement.DataContainer{}
	// nPost1 := datamanagement.Topics{}
	// nPost1.Title = "Affronte2"
	// nPost1.TopicID = 3
	// nPost1.CreationDate = time.Now()
	// nPost1.ValidTopic = true
	// nPost1.Upvotes = 998
	// nCTN.Topics = nPost1
	// datamanagement.AddLineIntoTargetTable(nCTN, "Topics")
	//handlers
	http.HandleFunc("/", handler.MainHandler)
	// allTopics := datamanagement.GetAllFromTable("Topics")
	// for _, t := range allTopics {
	// 	http.HandleFunc("/topic/"+t.Topics.Title, handler.Topic)
	// }
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("static/fonts"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("static/img"))))
	fmt.Println("(http://localhost"+port+"/home"+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}
