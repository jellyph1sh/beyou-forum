package datamanagement

import (
	"database/sql"
	"fmt"
	"strings"

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
	case table == "User":
		res, err = insertUserInUser.Exec(data.User.ID, data.User.Name, data.User.First_name, data.User.User_name, data.User.Email, data.User.Password, data.User.Is_admin, data.User.Is_valid, data.User.Description, data.User.Profile_image, data.User.Creation_date, strings.Join(strings.Fields(fmt.Sprint(data.User.Post_like)), ","), strings.Join(strings.Fields(fmt.Sprint(data.User.Post_dislike)), ","), strings.Join(strings.Fields(fmt.Sprint(data.User.Topic_like)), ","))
		break
	case table == "Post":
		res, err = insertUserInUser.Exec(data.Post.ID, data.Post.Like, data.Post.Author_id, data.Post.Is_valid, data.Post.Content, strings.Join(strings.Fields(fmt.Sprint(data.Post.Comentary)), ","), data.Post.Dislike, data.Post.Topic, data.Post.Date)
		break
	case table == "Topic":
		res, err = insertUserInUser.Exec(data.Topic.ID, data.Topic.Title, data.Topic.Description, data.Topic.Is_valid, strings.Join(strings.Fields(fmt.Sprint(data.Topic.Follow)), ","), data.Topic.Creator, data.Topic.Like)
		break
	case table == "Tag":
		res, err = insertUserInUser.Exec(data.Tag.ID, data.Tag.Title, data.Tag.Title)
		break
	case table == "Ban":
		res, err = insertUserInUser.Exec(data.Ban.ID, data.Ban.Word, data.Ban.Admin, data.Ban.Comment)
		break
	case table == "Report":
		res, err = insertUserInUser.Exec(data.Report.ID, data.Report.ID_post, data.Report.ID_post, data.Report.Comment)
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

func UpdateLike(table string, data DataContainer, add bool, id int) {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	defer db.Close()
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return
	}
	var updateLike *sql.Stmt
	if add {
		updateLike, err = db.Prepare("UPDATE " + table + " SET like=like+1 WHERE id = ?;")
	} else {
		updateLike, err = db.Prepare("UPDATE " + table + " SET dislike=like+1 WHERE id = ?;")
	}
	if err != nil {
		fmt.Println(err)
	}
	res, err := updateLike.Exec(id)
	affected, _ := res.RowsAffected()
	fmt.Println(affected, " ", table, " has got a new like/dislike")
}
