package common

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB(dbHost, dbPort, dbUser, dbPass, dbName string) (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println(connStr)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
