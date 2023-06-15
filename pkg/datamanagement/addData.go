package datamanagement

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func AddLineIntoTargetTable(data DataContainer, table string) {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return
	}
	defer db.Close()
	var insertDataInTable *sql.Stmt
	var res sql.Result
	switch true {
	case table == "Users":
		query := "INSERT INTO " + table + "(UserID, Username,Email,Password,Firstname,Lastname,Description,CreationDate,ProfilePicture,IsAdmin,ValidUser) VALUES (?,?,?,?,?,?,?,?,?,?,?);"
		insertDataInTable, err = db.Prepare(query)
		CheckPrepareQuery(err)
		res, err = insertDataInTable.Exec(data.Users.UserID, data.Users.Username, data.Users.Email, data.Users.Password, data.Users.Firstname, data.Users.Lastname, data.Users.Description, data.Users.CreationDate, data.Users.ProfilePicture, data.Users.IsAdmin, data.Users.ValidUser)
		break
	case table == "Posts":
		query := "INSERT INTO " + table + "(Content,AuthorID,TopicID,Likes,Dislikes,CreationDate,IsValidPost) VALUES(?,?,?,?,?,?,?);"
		insertDataInTable, err = db.Prepare(query)
		CheckPrepareQuery(err)
		res, err = insertDataInTable.Exec(data.Posts.Content, data.Posts.AuthorID, data.Posts.TopicID, data.Posts.Likes, data.Posts.Dislikes, data.Posts.CreationDate, data.Posts.IsValidPost)
		break
	case table == "Topics":
		query := "INSERT INTO " + table + "(Title,Description,Picture,CreationDate,CreatorID,Upvotes,Follows,ValidTopic) VALUES(?,?,?,?,?,?,?,?);"
		insertDataInTable, err = db.Prepare(query)
		CheckPrepareQuery(err)
		res, err = insertDataInTable.Exec(data.Topics.Title, data.Topics.Description, data.Topics.Picture, data.Topics.CreationDate, data.Topics.CreatorID, data.Topics.Upvotes, data.Topics.Follows, data.Topics.ValidTopic)
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
		query := "INSERT INTO " + table + "(PostID,ReportUserID,Comment) VALUES (?,?,?);"
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
	case table == "Upvotes":
		query := "INSERT INTO " + table + " VALUES (?,?);"
		insertDataInTable, err = db.Prepare(query)
		CheckPrepareQuery(err)
		res, err = insertDataInTable.Exec(data.Upvotes.TopicID, data.Upvotes.UserID)
		break
	default:
		fmt.Println("Invalid Table")
		return
	}
	if err != nil || res == nil {
		fmt.Println("Could not insert this data : \n", err)
		return
	}
	affected, _ := res.RowsAffected()
	fmt.Println(affected, " ", table, " has been add to the database")
}

func UpdateUpvotes(TopicID int, UserID string) {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	defer db.Close()
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return
	}
	var updateUpvotes *sql.Stmt
	sign := "+"
	row := ReadDB("SELECT * FROM Upvotes WHERE TopicID = " + strconv.Itoa(TopicID) + " AND UserID = " + UserID + ";")
	for row.Next() {
		sign = "-"
		DeleteLineIntoTargetTable("Upvotes", "TopicID = "+strconv.Itoa(TopicID)+" AND UserID = "+UserID)
		row.Close()
	}
	if sign == "+" {
		AddLineIntoTargetTable(DataContainer{Upvotes: Upvotes{TopicID: TopicID, UserID: UserID}}, "Upvotes")
	}
	updateUpvotes, err = db.Prepare("UPDATE Topics SET Upvotes=Upvotes" + sign + "1 WHERE TopicID = ?;")
	if err != nil {
		fmt.Println(err)
	}
	res, err := updateUpvotes.Exec(TopicID)
	affected, _ := res.RowsAffected()
	fmt.Println(affected, " upvote of upvotes/unupvotes")
}

/*
likOrdIS: 'Likes' - 'Dislikes';
*/
func LikePostManager(idPost int, idUser string, likOrdIS string) {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	defer db.Close()
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return
	}
	row := ReadDB("SELECT * FROM " + likOrdIS + " WHERE PostID = " + fmt.Sprint(idPost) + " AND UserID = " + fmt.Sprint(idUser) + ";")
	for row.Next() {
		row.Close()
		DeleteLineIntoTargetTable(likOrdIS, "PostID = "+fmt.Sprint(idPost)+" AND UserID = "+fmt.Sprint(idUser))
		updateLike, err := db.Prepare("UPDATE Posts SET " + likOrdIS + "=" + likOrdIS + "-1 WHERE PostID = ?;")
		if err != nil {
			fmt.Println(err)
		}
		updateLike.Exec(idPost)
		return
	}
	if likOrdIS == "Likes" {
		row := ReadDB("SELECT * FROM Dislikes WHERE PostID = " + fmt.Sprint(idPost) + " AND UserID = " + fmt.Sprint(idUser) + ";")
		for row.Next() {
			row.Close()
			DeleteLineIntoTargetTable("Dislikes", "PostID = "+fmt.Sprint(idPost)+" AND UserID = "+fmt.Sprint(idUser))
		}
		AddLineIntoTargetTable(DataContainer{Likes: Likes{PostID: idPost, UserID: idUser}}, "Likes")
	} else {
		row := ReadDB("SELECT * FROM Likes WHERE PostID = " + fmt.Sprint(idPost) + " AND UserID = " + fmt.Sprint(idUser) + ";")
		for row.Next() {
			row.Close()
			DeleteLineIntoTargetTable("Likes", "PostID = "+fmt.Sprint(idPost)+" AND UserID = "+fmt.Sprint(idUser))
		}
		AddLineIntoTargetTable(DataContainer{Dislikes: Dislikes{PostID: idPost, UserID: idUser}}, "Dislikes")
	}
	updateLike, err := db.Prepare("UPDATE Posts SET " + likOrdIS + "=" + likOrdIS + "+1 WHERE PostID = ?;")
	if err != nil {
		fmt.Println(err)
	}
	updateLike.Exec(idPost)
	return
}

func DeleteLineIntoTargetTable(table, condition string) {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	defer db.Close()
	if err != nil {
		fmt.Println("delete is not possible")
		fmt.Println("Could not open database : \n", err)
		return
	}
	query := "DELETE FROM " + table + " WHERE (" + condition + ");"
	fmt.Println(query)
	res, err := db.Exec(query)
	if err != nil {
		fmt.Println("delete is not possible")
		fmt.Println("invalid condition : \n", err)
		return
	}
	affected, _ := res.RowsAffected()
	fmt.Println(affected, "line of ", table, "has been deleted")
}

func UpdateLine(query string) {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	defer db.Close()
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return
	}
	db.Exec(query)
}
