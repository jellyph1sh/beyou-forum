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
	// topicTest := datamanagement.DataContainer{Topics: datamanagement.Topics{TopicID: 2, Title: "AffronteMonRegard", Description: "Siphano en sueur", Picture: "https://image.noelshack.com/fichiers/2019/34/4/1566485216-ecmspsewkaegaxk.jpg", CreatorID: "a1fce4363854ff888cff4b8e7875d600c2682390412a8cf79b37d0b11148b0fa", CreationDate: time.Now(), Upvotes: 32, Follows: 2, ValidTopic: true}}
	// postTest := datamanagement.DataContainer{Posts: datamanagement.Posts{Content: "un autre post test odiajdiuoazudanjudfnajudfnaefn", AuthorID: "a1fce4363854ff888cff4b8e7875d600c2682390412a8cf79b37d0b11148b0fa", TopicID: 2, PostID: 2, Likes: 15, Dislikes: 0, CreationDate: time.Now(), IsValidPost: true}}
	// // userTest := datamanagement.DataContainer{Users: datamanagement.Users{UserID: "1", Username: "DarkSasuke", Email: "email", Password: "oipjfziiofnziofnez", Firstname: "Dark", Lastname: "Link", Description: "baka", CreationDate: time.Now(), ProfilePicture: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQLeR3wX3QrJpZGlb6fIeU-XPPRgxGlP5coYmQaBFZJ&s", IsAdmin: false, ValidUser: true}}
	// datamanagement.AddLineIntoTargetTable(postTest, "Posts")
	// datamanagement.UpdateLine("UPDATE Topics SET Title='test' WHERE TopicID = 2;")
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
