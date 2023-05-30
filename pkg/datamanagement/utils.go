package datamanagement

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	username = ""
	password = ""
	hostname = ""
	dbname   = ""
)

type Display struct {
	Logout bool
}
type User struct {
	id            int
	name          string
	first_name    string
	user_name     string
	email         string
	password      string
	is_admin      bool
	is_valid      bool
	description   string
	profile_image string
	creation_date time.Time
}

type post struct {
	id        int
	like      []int
	author_id int
	is_valid  bool
	content   string
	comentary []int
	dislike   []int
	topic     int
}

type topic struct {
	id          int
	title       string
	description string
	is_valid    bool
	follow      []int
	creator     int
}

type tag struct {
	id      int
	title   string
	creator int
	like    []int
}

/*don't forget to close the *sql.Rows when you use this func */
func readDB(table, query string) *sql.Rows {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	defer db.Close()
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return nil
	}
	row, err := db.Query("SELECT * FROM User;")
	if err != nil {
		fmt.Println("Invalid request :")
		log.Fatal(err)
		return nil
	}
	return row
}
