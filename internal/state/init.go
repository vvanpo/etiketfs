package state

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Init(db *sql.DB) (State, error) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS file (
			id TEXT PRIMARY KEY NOT NULL,
			added INTEGER NOT NULL
		);
	`)

	if err == nil {
		return State{db}, nil
	}

	return State{}, err
}
