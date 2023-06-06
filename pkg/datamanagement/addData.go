package datamanagement

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func AddLineIntoTargetTable(data DataContainer, table string) {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	defer db.Close()
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return
	}
	var insertDataInTable *sql.Stmt
	var res sql.Result
	switch true {
	case table == "Users":
		query := "INSERT INTO " + table + "(UserName,Email,Password,Firstname,Lastname,Description,CreationDate,ProfilePicture,IsAdmin,ValidUser) VALUES(?,?,?,?,?,?,?,?,?,?);"
		insertDataInTable, err = db.Prepare(query)
		CheckPrepareQuery(err)
		res, err = insertDataInTable.Exec(data.Users.Username, data.Users.Email, data.Users.Password, data.Users.Firstname, data.Users.Lastname, data.Users.Description, data.Users.CreationDate, data.Users.ProfilePicture, data.Users.IsAdmin, data.Users.ValidUser)
		break
	case table == "Posts":
		query := "INSERT INTO " + table + "(Content,AuthorID,TopicID,Likes,Dislikes,CreationDate,IsValidPost) VALUES(?,?,?,?,?,?,?);"
		insertDataInTable, err = db.Prepare(query)
		CheckPrepareQuery(err)
		res, err = insertDataInTable.Exec(data.Posts.Content, data.Posts.AuthorID, data.Posts.TopicID, data.Posts.Likes, data.Posts.Dislikes, data.Posts.CreationDate, data.Posts.IsValidPost)
		break
	case table == "Topics":
		query := "INSERT INTO " + table + "(Title,Description,CreationDate,CreatorID,Upvotes,Follows,ValidTopic) VALUES(?,?,?,?,?,?);"
		insertDataInTable, err = db.Prepare(query)
		CheckPrepareQuery(err)
		res, err = insertDataInTable.Exec(data.Topics.Title, data.Topics.Description, data.Topics.CreationDate, data.Topics.CreatorID, data.Topics.Upvotes, data.Topics.Follows, data.Topics.ValidTopic)
		break
	case table == "Tags":
		query := "INSERT INTO " + table + "(Title,CreatorID) VALUES(?,?);"
		insertDataInTable, err = db.Prepare(query)
		CheckPrepareQuery(err)
		res, err = insertDataInTable.Exec(data.Tags.Title, data.Tags.CreatorID)
		break
	case table == "WordsBlacklist":
		query := "INSERT INTO " + table + "(word) VALUES(?);"
		insertDataInTable, err = db.Prepare(query)
		CheckPrepareQuery(err)
		res, err = insertDataInTable.Exec(data.WordsBlacklist.Word)
		break
	case table == "Reports":
		query := "INSERT INTO " + table + "(PostID,ReportUserID,Comment) VALUES (?,?,?,?);"
		insertDataInTable, err = db.Prepare(query)
		CheckPrepareQuery(err)
		res, err = insertDataInTable.Exec(data.Reports.PostID, data.Reports.ReportUserID, data.Reports.Comment)
		break
	case table == "Follows":
		query := "INSERT INTO " + table + "(TopicID,UserID) VALUES (?,?);"
		insertDataInTable, err = db.Prepare(query)
		CheckPrepareQuery(err)
		res, err = insertDataInTable.Exec(data.Follows.TopicID, data.Follows.UserID)
		break
	case table == "Likes":
		query := "INSERT INTO " + table + " VALUES (?,?);"
		insertDataInTable, err = db.Prepare(query)
		CheckPrepareQuery(err)
		res, err = insertDataInTable.Exec(data.Likes.PostID, data.Likes.UserID)
		break
	case table == "Dislikes":
		query := "INSERT INTO " + table + " VALUES (?,?);"
		insertDataInTable, err = db.Prepare(query)
		CheckPrepareQuery(err)
		res, err = insertDataInTable.Exec(data.Dislikes.PostID, data.Dislikes.UserID)
		break
	case table == "TopicsTags":
		query := "INSERT INTO " + table + " VALUES (?,?);"
		insertDataInTable, err = db.Prepare(query)
		CheckPrepareQuery(err)
		res, err = insertDataInTable.Exec(data.TopicsTags.TopicID, data.TopicsTags.TagID)
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

func UpdateUpvotes(table string, data DataContainer, id int) {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	defer db.Close()
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return
	}
	var updateUpvotes *sql.Stmt
	updateUpvotes, err = db.Prepare("UPDATE Topic SET Upvotes=Upvotes+1 WHERE id = ?;")
	if err != nil {
		fmt.Println(err)
	}
	res, err := updateUpvotes.Exec(id)
	affected, _ := res.RowsAffected()
	fmt.Println(affected, " ", table, " has got a new upvotes/unupvotes")
}
