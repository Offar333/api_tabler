package controllers

import (
	"database/sql"
)

var db *sql.DB
var err error
var connStr string

func main() {

	connStr = "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db"

}
