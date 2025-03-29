package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "go_crud"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:8889)/"+dbName)
	return db, err
}
