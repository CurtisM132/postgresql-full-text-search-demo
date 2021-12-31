package search

import "database/sql"

const connStr = "postgres://postgres:abc123@localhost:4200/shakespeare_plays?sslmode=disable"

func connectToDB() (*sql.DB, error) {
	return sql.Open("postgres", connStr)
}
