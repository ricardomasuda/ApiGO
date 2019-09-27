package DataBase

import (
	"database/sql"
	// Driver Mysql para Go
	_ "github.com/go-sql-driver/mysql"
	)

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "crudnode"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic("[DataBse] Erro de banco"+err.Error())
	}
	return db
}
