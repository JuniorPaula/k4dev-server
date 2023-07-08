package database

import (
	"database/sql"
	"fmt"
	"knowledge-api/internal/config"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", config.ConnectStringPostgres)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}

func ConnectToDB() (*sql.DB, error) {
	conn, err := openDB()
	if err != nil {
		return nil, err
	}
	fmt.Println("[::] Connected to Postgres database")

	return conn, nil
}
