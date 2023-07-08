package postgres

import "database/sql"

func Connect() (*sql.DB, error) {
	return sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/knowledge?sslmode=disable")
}
