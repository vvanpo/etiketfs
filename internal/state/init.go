package state

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func Init(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS file (
			id TEXT PRIMARY KEY NOT NULL,
			added INTEGER NOT NULL
		);
	`)

	return err
}

func AddFile(db *sql.DB, id uuid.UUID) error {
	ts := time.Now().Unix()
	_, err := db.Exec(`INSERT INTO file VALUES (?, ?);`, id, ts)

	return err
}
