package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Getlink() *sql.DB {
	db, err := sql.Open("mysql", "root:gheao126@tcp(127.0.0.1)/userdata")
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	return db
}
