package main

import (
	"fmt"
	"forum/pkg/handler"
	"net/http"
)

var port = ":8080"

func main() {
	//handlers
	// http.HandleFunc("/account", handler.Account)
	// http.HandleFunc("/automod", handler.Automod)
	// http.HandleFunc("/explore", handler.Explore)
	// http.HandleFunc("/home", handler.Home)
	// http.HandleFunc("/login", handler.Login)
	// http.HandleFunc("/profile", handler.Profile)
	// http.HandleFunc("/register", handler.Register)
	// http.HandleFunc("/topic", handler.Topic)

	http.HandleFunc("/", handler.MainHandler)
	fmt.Println("(http://localhost"+port+"/home"+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}
