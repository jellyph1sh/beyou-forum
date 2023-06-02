package datamanagement

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func IsUserExist(userInput string, password string) bool {
	passwordByte := []byte(password)
	passwordInSha256 := sha256.Sum256(passwordByte)
	stringPasswordInSha256 := fmt.Sprintf("%x", passwordInSha256[:])
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return false
	}
	defer db.Close()
	row, err := db.Query("SELECT * FROM Users;")
	if err != nil {
		fmt.Println("Invalid request :")
		log.Fatal(err)
		return false
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var UserID int
		var Username string
		var Email string
		var Password string
		var Firstname string
		var Lastname string
		var Description string
		var CreationDate time.Time
		var ProfilePicture string
		var IsAdmin bool
		var ValidUser bool
		row.Scan(&UserID, &Username, &Email, &Password, &Firstname, &Lastname, &Description, &CreationDate, &ProfilePicture, &IsAdmin, &ValidUser)
		if (Lastname == userInput || Email == userInput) && Password == stringPasswordInSha256 {
			return true
		}
	}
	return false
}
