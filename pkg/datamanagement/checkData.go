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
	row, err := db.Query("SELECT * FROM User;")
	if err != nil {
		fmt.Println("Invalid request :")
		log.Fatal(err)
		return false
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
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
		row.Scan(&id, &name, &first_name, &user_name, &email, &password, &is_admin, &is_valid, &description, &profile_image, &creation_date)
		if (name == userInput || email == userInput) && password == stringPasswordInSha256 {
			return true
		}
	}
	return false
}
