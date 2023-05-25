package datamanagement

import (
	"database/sql"
	"fmt"
	"reflect"
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
	case !reflect.ValueOf(data.User).IsZero():
		res, err = insertUserInUser.Exec(data.User.ID, data.User.Name, data.User.First_name, data.User.User_name, data.User.Email, data.User.Password, data.User.Is_admin, data.User.Is_valid, data.User.Description, data.User.Profile_image, data.User.Creation_date)
		break
	case !reflect.ValueOf(data.Post).IsZero():
		res, err = insertUserInUser.Exec(data.Post.ID, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data.Post.Like)), ","), "[]"), data.Post.Author_id, data.Post.Is_valid, data.Post.Content, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data.Post.Comentary)), ","), "[]"), strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data.Post.Dislike)), ","), "[]"), data.Post.Topic)
		break
	case !reflect.ValueOf(data.Topic).IsZero():
		res, err = insertUserInUser.Exec(data.Topic.ID, data.Topic.Title, data.Topic.Description, data.Topic.Is_valid, data.Topic.Follow, data.Topic.Creator)
		break
	case !reflect.ValueOf(data.Tag).IsZero():
		res, err = insertUserInUser.Exec(data.Topic.ID, data.Topic.Title, data.Topic.Description, data.Topic.Is_valid, data.Topic.Follow, data.Topic.Creator)
		break
	case !reflect.ValueOf(data.Ban).IsZero():
		res, err = insertUserInUser.Exec(data.Ban.ID, data.Ban.Word, data.Ban.Admin, data.Ban.Comment)
		break
	case !reflect.ValueOf(data.Report).IsZero():
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
	fmt.Println(affected, " user has been add to the database")
}
