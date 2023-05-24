package datamanagement

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = ""
	password = ""
	hostname = ""
	dbname   = ""
)

type user struct {
	id            int
	name          string
	first_name    string
	user_name     string
	email         string
	password      string
	is_admin      bool
	is_valid      bool
	description   string
	profile_image string
	creation_date time.Time
}

type post struct {
	id        int
	like      []int
	author_id int
	is_valid  bool
	content   string
	comentary []int
	dislike   []int
	topic     int
}

type topic struct {
	id          int
	title       string
	description string
	is_valid    bool
	follow      []int
	creator     int
}

type tag struct {
	id      int
	title   string
	creator int
	like    []int
}

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}
