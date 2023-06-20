package datamanagement

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func IsPostDLikeByBYser(PostID int, UserID string, DisOrLike string) bool {
	rows := SelectDB("SELECT * FROM ? WHERE PostID = ? AND UserID = ?", DisOrLike, strconv.Itoa(PostID), UserID)
	defer rows.Close()

	for rows.Next() {
		return true
	}
	return false
}

func IsUserExist(userEmail string, userProfilName string) bool {
	rows := SelectDB("SELECT * FROM Users WHERE ( Username = ? OR Email = ?);", string(userProfilName), string(userEmail))
	defer rows.Close()

	if !rows.Next() {
		return false
	} else {
		return true
	}
}

func IsRegister(userInput string, password string) (bool, string) {
	passwordByte := []byte(password)
	passwordInSha256 := sha256.Sum256(passwordByte)
	stringPasswordInSha256 := fmt.Sprintf("%x", passwordInSha256[:])
	rows := SelectDB("SELECT userID FROM Users WHERE ( Username = ? OR Email = ? ) AND Password = ?;", string(userInput), string(userInput), string(stringPasswordInSha256))
	defer rows.Close()

	for rows.Next() {
		var id string
		rows.Scan(&id)
		return true, id
	}

	return false, ""
}
