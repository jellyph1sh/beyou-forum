package datamanagement

import (
	"datamanagement/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func isUserExist(userName, email string) {
	db, err := sql.Open("mysql", dsn())
	if err != nil {
		fmt.Println("Could not open database : \n", err)
		return
	}
	defer db.Close()
	
}
