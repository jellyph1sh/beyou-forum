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
	ID            int
	Name          string
	First_name    string
	User_name     string
	Email         string
	Password      string
	Is_admin      bool
	Is_valid      bool
	Description   string
	Profile_image string
	Creation_date time.Time
}

type Post struct {
	ID        int
	Like      []int
	Author_id int
	Is_valid  bool
	Content   string
	Comentary []int
	Dislike   []int
	Topic     int
}

type Topic struct {
	ID          int
	Title       string
	Description string
	Is_valid    bool
	Follow      []int
	Creator     int
}

type Tag struct {
	ID      int
	Title   string
	Creator int
	Like    []int
}

type Ban struct {
	ID      int
	Word    string
	Admin   int
	Comment string
}

type Report struct {
	ID      int
	ID_post int
	ID_user int
	Comment string
}

type DataContainer struct {
	User   User
	Post   Post
	Topic  Topic
	Tag    Tag
	Ban    Ban
	Report Report
}

type DataForOnePost struct {
	NbLike     int
	AuthorName string
	Is_valid   bool
	Content    string
	Comentary  []int
	NBDislike  int
	TopicName  string
}

/*don't forget to close the *sql.Rows when you use this func */
func readDB(query string) *sql.Rows {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	defer db.Close()
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return nil
	}
	row, err := db.Query(query)
	if err != nil {
		fmt.Println("Invalid request :")
		log.Fatal(err)
		return nil
	}
	return row
}

func buildQueryAddData(table string, nbValues int) string {
	result := "INSERT INTO " + table + " Values (?"
	for i := 1; i < nbValues; i++ {
		result += ",?"
	}
	return result + ");"
}
