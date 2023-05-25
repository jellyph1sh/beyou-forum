package datamanagement

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func IsUserExist(userName, email string) {
	db, err := sql.Open("sqlite3", "./DB-Forum.db")
	defer db.Close()
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return
	}
	row, err := db.Query("SELECT * FROM User;")
	if err != nil {
		fmt.Println("Invalid request :")
		log.Fatal(err)
		return
	}
	defer row.Close()
	// fmt.Println(row)
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
		fmt.Println("Student: ", name)
	}
}
