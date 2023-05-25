package main

import (
	"fmt"
	"forum/pkg/datamanagement"
	"forum/pkg/handler"
	"net/http"
)

var port = ":8080"

func main() {
	datamanagement.IsUserExist("", "")
	//handlers
	http.HandleFunc("/", handler.MainHandler)
	fmt.Println("(http://localhost"+port+"/home"+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}
