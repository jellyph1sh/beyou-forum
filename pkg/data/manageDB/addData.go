package datamanagement

import (
	"datamanagement/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func AddUser(userData user) {
	db, err := sql.Open("mysql", dsn())
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return
	}
	defer db.Close()
	// TO DO : check user data
	query := "INSERT INTO 'users' (id,name,first_name,user_name,email,password,is_admin,is_valid,description,profile_image,creation_date) VALUES (?,?,?,?,?,?,?,?,?,?,?);"
	insert, err := db.Prepare(query)
	if err != nil {
		fmt.Println("Error when try to make insert query : \n", err)
		return
	}
	resp, err := insert.Exec(userData.id, userData.name, userData.first_name, userData.user_name, userData.email, userData.password, userData.is_admin, userData.is_valid, userData.description, userData.profile_image, userData.creation_date)
	insert.Close()
	if err != nil {
		fmt.Println("Could not insert this data : \n", userData, "\n", err)
		return
	}
}
