package datamanagement

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func AddUser() {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	defer db.Close()
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return
	}
	// TO DO : check user data
	query := "INSERT INTO 'User' (ID,name,first_name,user_name,email,password,is_admin,is_valid,description,profile_img,creation_date) VALUES (?,?,?,?,?,?,?,?,?,?,?);"
	insertUserUser, err := db.Prepare(query)
	if err != nil {
		fmt.Println("Could not prepare request :", "\n", err)
		return
	}
	res, err := insertUserUser.Exec(query)
	if err != nil {
		fmt.Println("Could not insert this data : \n", "\n", err)
		return
	}
	fmt.Println(res)
}
