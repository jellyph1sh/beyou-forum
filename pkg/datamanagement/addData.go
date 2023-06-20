package datamanagement

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func AddLineIntoTargetTable(data DataContainer, table string) {
	var res sql.Result
	switch true {
	case table == "Users":
		res = AddDeleteUpdateDB("INSERT INTO Users (UserID, Username,Email,Password,Firstname,Lastname,Description,CreationDate,ProfilePicture,IsAdmin,ValidUser) VALUES (?,?,?,?,?,?,?,?,?,?,?);", data.Users.UserID, data.Users.Username, data.Users.Email, data.Users.Password, data.Users.Firstname, data.Users.Lastname, data.Users.Description, data.Users.CreationDate, data.Users.ProfilePicture, data.Users.IsAdmin, data.Users.ValidUser)
		break
	case table == "Posts":
		res = AddDeleteUpdateDB("INSERT INTO Posts (Content,AuthorID,TopicID,Likes,Dislikes,CreationDate,IsValidPost) VALUES(?,?,?,?,?,?,?);", data.Posts.Content, data.Posts.AuthorID, data.Posts.TopicID, data.Posts.Likes, data.Posts.Dislikes, data.Posts.CreationDate, data.Posts.IsValidPost)
		break
	case table == "Topics":
		res = AddDeleteUpdateDB("INSERT INTO Topics (Title,Description,Picture,CreationDate,CreatorID,Upvotes,Follows,ValidTopic) VALUES(?,?,?,?,?,?,?,?);", data.Topics.Title, data.Topics.Description, data.Topics.Picture, data.Topics.CreationDate, data.Topics.CreatorID, data.Topics.Upvotes, data.Topics.Follows, data.Topics.ValidTopic)
		break
	case table == "Tags":
		res = AddDeleteUpdateDB("INSERT INTO Tags (Title,CreatorID) VALUES(?,?);", data.Tags.Title, data.Tags.CreatorID)
		break
	case table == "WordsBlacklist":
		res = AddDeleteUpdateDB("INSERT INTO WordsBlacklist (word) VALUES(?);", data.WordsBlacklist.Word)
		break
	case table == "Reports":
		res = AddDeleteUpdateDB("INSERT INTO Reports (PostID,ReportUserID,Comment) VALUES (?,?,?);", data.Reports.PostID, data.Reports.ReportUserID, data.Reports.Comment)
		break
	case table == "Follows":
		res = AddDeleteUpdateDB("INSERT INTO Follows (TopicID,UserID) VALUES (?,?);", data.Follows.TopicID, data.Follows.UserID)
		break
	case table == "Likes":
		res = AddDeleteUpdateDB("INSERT INTO Likes VALUES (?,?);", data.Likes.PostID, data.Likes.UserID)
		break
	case table == "Dislikes":
		res = AddDeleteUpdateDB("INSERT INTO Dislikes VALUES (?,?);", data.Dislikes.PostID, data.Dislikes.UserID)
		break
	case table == "TopicsTags":
		res = AddDeleteUpdateDB("INSERT INTO TopicsTags VALUES (?,?);", data.TopicsTags.TopicID, data.TopicsTags.TagID)
		break
	case table == "Upvotes":
		res = AddDeleteUpdateDB("INSERT INTO Upvotes VALUES (?,?);", data.Upvotes.TopicID, data.Upvotes.UserID)
		break
	default:
		fmt.Println("Invalid Table")
		return
	}
	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(affected, " ", table, " has been add to the database")
}

func UpdateUpvotes(TopicID int, UserID string) {
	sign := "+"
	rows := SelectDB("SELECT * FROM Upvotes WHERE TopicID = ? AND UserID = ?;", strconv.Itoa(TopicID), UserID)
	defer rows.Close()
	for rows.Next() {
		sign = "-"
		res := AddDeleteUpdateDB("DELETE FROM Upvotes WHERE TopicID = ? AND UserID = ?;", strconv.Itoa(TopicID), UserID)
		affected, err := res.RowsAffected()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(affected, "deleted!")
	}
	if sign == "+" {
		AddLineIntoTargetTable(DataContainer{Upvotes: Upvotes{TopicID: TopicID, UserID: UserID}}, "Upvotes")
	}

	res := AddDeleteUpdateDB("UPDATE Topics SET Upvotes=Upvotes ?1 WHERE TopicID = ?;", sign, TopicID)
	affected, _ := res.RowsAffected()
	fmt.Println(affected, " upvote of upvotes/unupvotes")
}

/*
likOrdIS: 'Likes' - 'Dislikes';
*/
func LikePostManager(idPost int, idUser string, likOrdIS string) {
	rows := SelectDB("SELECT * FROM ? WHERE PostID = ? AND UserID = ?;", likOrdIS, fmt.Sprint(idPost), fmt.Sprint(idUser))
	defer rows.Close()
	for rows.Next() {
		res := AddDeleteUpdateDB("DELETE FROM ? WHERE PostID = ? AND UserID = ?;", likOrdIS, fmt.Sprint(idPost), fmt.Sprint(idUser))
		affected, err := res.RowsAffected()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(affected, "deleted!")
		res = AddDeleteUpdateDB("UPDATE Posts SET ? = ?-1 WHERE PostID = ?;", likOrdIS, likOrdIS, idPost)
		affected, err = res.RowsAffected()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(affected, "updated!")
		return
	}
	if likOrdIS == "Likes" {
		rows := SelectDB("SELECT * FROM Dislikes WHERE PostID = ? AND UserID = ?;", fmt.Sprint(idPost), fmt.Sprint(idUser))
		defer rows.Close()
		for rows.Next() {
			res := AddDeleteUpdateDB("DELETE FROM Dislikes WHERE PostID = ? AND UserID = ?;", fmt.Sprint(idPost), fmt.Sprint(idUser))
			affected, err := res.RowsAffected()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(affected, "deleted!")
		}
		AddLineIntoTargetTable(DataContainer{Likes: Likes{PostID: idPost, UserID: idUser}}, "Likes")
	} else {
		rows := SelectDB("SELECT * FROM Likes WHERE PostID = ? AND UserID = ?;", fmt.Sprint(idPost), fmt.Sprint(idUser))
		rows.Close()
		for rows.Next() {
			res := AddDeleteUpdateDB("DELETE FROM Likes WHERE PostID = ? AND UserID = ?;", fmt.Sprint(idPost), fmt.Sprint(idUser))
			affected, err := res.RowsAffected()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(affected, "deleted!")
		}
		AddLineIntoTargetTable(DataContainer{Dislikes: Dislikes{PostID: idPost, UserID: idUser}}, "Dislikes")
	}
	res := AddDeleteUpdateDB("UPDATE Posts SET ? = ?-1 WHERE PostID = ?;", likOrdIS, likOrdIS, idPost)
	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(affected, "updated!")
	return
}

func AddTagsToTopic(tags, creatorId string, TopicID int) {
	tagsArray := strings.Split(tags, " ")
	for _, tag := range tagsArray {
		if (GetTagByName(tag) == Tags{}) {
			AddLineIntoTargetTable(DataContainer{Tags: Tags{Title: tag, CreatorID: creatorId}}, "Tags")
		}
		AddLineIntoTargetTable(DataContainer{TopicsTags: TopicsTags{TopicID: TopicID, TagID: GetTagByName(tag).TagID}}, "TopicsTags")
	}
}
