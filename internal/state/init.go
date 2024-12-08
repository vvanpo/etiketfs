package state

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Init(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE file (
			id TEXT PRIMARY KEY NOT NULL
		);
	`)

	return err
}
