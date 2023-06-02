package datamanagement

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func AddLineIntoTargetTable(data DataContainer, table string, nbValues int) {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	defer db.Close()
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return
	}

	insertUserInUser, err := db.Prepare(buildQueryAddData(table, nbValues))
	if err != nil {
		fmt.Println("Could not prepare request :", "\n", err)
		return
	}
	var res sql.Result
	switch true {
	case table == "Users":
		res, err = insertUserInUser.Exec(data.Users.UserID, data.Users.Username, data.Users.Email, data.Users.Password, data.Users.Firstname, data.Users.Lastname, data.Users.Description, data.Users.CreationDate, data.Users.ProfilePicture, data.Users.IsAdmin, data.Users.ValidUser)
		break
	case table == "Posts":
		res, err = insertUserInUser.Exec(data.Posts.PostID, data.Posts.Content, data.Posts.AuthorID, data.Posts.TopicID, data.Posts.Likes, data.Posts.Dislikes, data.Posts.CreationDate, data.Posts.IsValidPost)
		break
	case table == "Topics":
		res, err = insertUserInUser.Exec(data.Topics.TopicID, data.Topics.Title, data.Topics.Description, data.Topics.CreatorID, data.Topics.Upvotes, data.Topics.Follows, data.Topics.ValidTopic)
		break
	case table == "Tags":
		res, err = insertUserInUser.Exec(data.Tags.TagID, data.Tags.Title, data.Tags.CreatorID)
		break
	case table == "WordsBlacklist":
		res, err = insertUserInUser.Exec(data.WordsBlacklist.WordID, data.WordsBlacklist.Word)
		break
	case table == "Reports":
		res, err = insertUserInUser.Exec(data.Reports.ReportID, data.Reports.PostID, data.Reports.ReportUserID, data.Reports.Comment)
		break
	default:
		fmt.Println("No data to add")
		return
	}
	if err != nil || res == nil {
		fmt.Println("Could not insert this data : \n", "\n", err)
		return
	}
	affected, _ := res.RowsAffected()
	fmt.Println(affected, " ", table, " has been add to the database")
}

func UpdateUpvotes(table string, data DataContainer, add bool, id int) {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	defer db.Close()
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return
	}
	var updateUpvotes *sql.Stmt
	if add || table == "Topics" {
		updateUpvotes, err = db.Prepare("UPDATE " + table + " SET Upvotes=Upvotes+1 WHERE id = ?;")
	} else {
		updateUpvotes, err = db.Prepare("UPDATE " + table + " SET Upvotes=Upvotes-1 WHERE id = ?;")
	}
	if err != nil {
		fmt.Println(err)
	}
	res, err := updateUpvotes.Exec(id)
	affected, _ := res.RowsAffected()
	fmt.Println(affected, " ", table, " has got a new upvotes/unupvotes")
}
