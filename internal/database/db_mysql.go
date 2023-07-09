package database

import (
	"database/sql"
	"knowledge-api/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect_MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectStringMySQL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
