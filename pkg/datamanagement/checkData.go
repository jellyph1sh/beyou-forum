package datamanagement

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func IsEmailAlreadyExist(userEmail string) bool {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Users WHERE (Email = ?);", string(userEmail))
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer rows.Close()

	if !rows.Next() {
		return false
	} else {
		return true
	}
}

func IsUsernameAlreadyExist(userProfilName string) bool {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Users WHERE ( Username = ?);", string(userProfilName))
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer rows.Close()

	if !rows.Next() {
		return false
	} else {
		return true
	}
}

func IsValidTopic(topic string) bool {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Topics WHERE (Title = ?);", string(topic))
	if err != nil {
		log.Fatal(err)
		return false
	}
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
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return false, ""
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Users WHERE ( Username = ? OR Email = ? ) AND Password = ?;", string(userInput), string(userInput), string(stringPasswordInSha256))
	if err != nil {
		log.Fatal(err)
		return false, "2"
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		var first_name string
		var user_name string
		var email string
		var password string
		var is_admin bool
		var is_valid bool
		var description string
		var profile_image string
		var creation_date time.Time
		rows.Scan(&id, &name, &first_name, &user_name, &email, &password, &is_admin, &is_valid, &description, &profile_image, &creation_date)
		return true, id
	}
	return false, "3"
}

func IsValidPassword(password string, idUser string) bool {
	dataUser := GetUserByID(idUser)
	passwordByte := []byte(password)
	passwordInSha256 := sha256.Sum256(passwordByte)
	stringPasswordInSha256 := fmt.Sprintf("%x", passwordInSha256[:])
	if dataUser.Password == stringPasswordInSha256 {
		return true
	} else {
		return false
	}
}

func IsLikeByUser(userID string, postID int) bool {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Likes WHERE (PostID = '?' AND UserID = ?);", userID, postID)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer rows.Close()

	if !rows.Next() {
		return false
	} else {
		return true
	}
}

func IsDislikeByUser(userID string, postID int) bool {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Dislikes WHERE (PostID = '?' AND UserID = ?);", userID, postID)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer rows.Close()

	if !rows.Next() {
		return false
	} else {
		return true
	}
}
