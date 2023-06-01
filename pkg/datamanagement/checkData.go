package datamanagement

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func IsUserExist(userEmail string, userProfilName string) bool {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return false
	}
	defer db.Close()
	QUERY := "SELECT * FROM User WHERE ( user_name = '" + string(userProfilName) + "' OR email = '" + string(userEmail) + "');"
	row, err := db.Query(QUERY)
	if err != nil {
		fmt.Println("Invalid request :")
		log.Fatal(err)
		return false
	}
	defer row.Close()
	if row.Next() == false {
		return false
	} else {
		return true
	}
}

func IsRegister(userInput string, password string) (bool, int) {
	passwordByte := []byte(password)
	passwordInSha256 := sha256.Sum256(passwordByte)
	stringPasswordInSha256 := fmt.Sprintf("%x", passwordInSha256[:])
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return false, 1
	}
	defer db.Close()
	QUERY := "SELECT * FROM User WHERE ( user_name = '" + string(userInput) + "' OR email = '" + string(userInput) + "' ) AND password = '" + string(stringPasswordInSha256) + "';"
	row, err := db.Query(QUERY)
	if err != nil {
		fmt.Println("Invalid request :")
		log.Fatal(err)
		return false, 1
	}
	defer row.Close()
	if row.Next() == false {
		return false, 1
	} else {
		for row.Next() {
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
			return true, id
		}
	}
	return false, 1
}
