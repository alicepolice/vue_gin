package utils

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"wolflong.com/vue_gin/variable"
)

func MySqlDB() {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "sql2008"
	dbName := "vue"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(256)
	db.SetMaxIdleConns(256)
	variable.DB = db
}
