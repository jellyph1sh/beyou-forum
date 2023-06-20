package datamanagement

import (
	"database/sql"
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
	UserID string
}

type Follows struct {
	FollowID int
	TopicID  int
	UserID   string
}

type Likes struct {
	PostID int
	UserID string
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
	TopicID      int
	Title        string
	Description  string
	Picture      string
	CreationDate time.Time
	CreatorID    string
	Upvotes      int
	Follows      int
	ValidTopic   bool
}

type TopicsTags struct {
	TopicID int
	TagID   int
}

type Upvotes struct {
	TopicID int
	UserID  string
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
	Topic    Topics
	Posts    []Posts
	Authors  []Users
	IsFollow bool
	IsUpvote bool
	Likes    []bool
	Dislikes []bool
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

func SelectDB(query string, args ...interface{}) *sql.Rows {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return rows
}

func AddDeleteUpdateDB(query string, args ...interface{}) sql.Result {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

	res, err := db.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return res
}
