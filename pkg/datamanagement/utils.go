package datamanagement

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	username = ""
	password = ""
	hostname = ""
	dbname   = ""
)

type DataFilter struct {
	number int
}

type Display struct {
	Logout bool
}

type Dislikes struct {
	PostID int
	UserID int
}

type Follows struct {
	FollowID int
	TopicID  int
	UserID   int
}

type Likes struct {
	PostID int
	UserID int
}

type Posts struct {
	PostID       int
	Content      string
	AuthorID     string
	TopicID      int
	Likes        int
	Dislikes     int
	CreationDate time.Time
	IsValidPost  bool
}

type Reports struct {
	ReportID     int
	PostID       int
	ReportUserID int
	Comment      string
}

type Tags struct {
	TagID     int
	Title     string
	CreatorID int
}

type UserConnected struct {
	IsUserConnected bool
	IdUser          string
}

type Topics struct {
	TopicID     int
	Title       string
	Description string
	Picture     string
	CreatorID   int
	Upvotes     int
	Follows     int
	ValidTopic  bool
}

type TopicsTags struct {
	TopicID int
	TagID   int
}

type Upvotes struct {
	TopicID int
	UserID  int
}

type Users struct {
	UserID         string
	Username       string
	Email          string
	Password       string
	Firstname      string
	Lastname       string
	Description    string
	CreationDate   time.Time
	ProfilePicture string
	IsAdmin        bool
	ValidUser      bool
}

type WordsBlacklist struct {
	WordID int
	Word   string
}

type DataContainer struct {
	Dislikes   Dislikes
	Follows    Follows
	Likes      Likes
	Posts      Posts
	Reports    Reports
	Tags       Tags
	Topics     Topics
	TopicsTags TopicsTags
	Upvotes    Upvotes
	Users
	WordsBlacklist
}

type DataTopicPage struct {
	Topic   Topics
	Posts   []Posts
	Authors []Users
}

/*don't forget to close the *sql.Rows when you use this func */
func ReadDB(query string) *sql.Rows {
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

func CheckPrepareQuery(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ExecuterQuery(QUERY string) {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		fmt.Println("Could not open database : \n", err)
	}
	defer db.Close()
	db.Exec(QUERY)
}

func CheckContentByBlackListWord(content string) bool {
	blackListWords := GetAllFromTable("WordsBlacklist")
	contentArray := strings.Split(content, " ")
	for _, w := range blackListWords {
		if arrayContains(contentArray, w.WordsBlacklist.Word) {
			return false
		}
	}

	return true
}

func arrayContains(array []string, word string) bool {
	for _, val := range array {
		if word == val {
			return true
		}
	}
	return false
}
