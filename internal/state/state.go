package state

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type State struct {
	db *sql.DB
}

func (s State) AddFile(id uuid.UUID) error {
	ts := time.Now().Unix()
	_, err := s.db.Exec(`INSERT INTO file VALUES (?, ?);`, id, ts)

	return err
}

// FileIds ...
func (s State) FileIds() (ids []uuid.UUID, err error) {
	rows, err := s.db.Query(`SELECT id FROM file`)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var nu uuid.NullUUID

		if err = rows.Scan(&nu); err != nil {
			return
		}

		if !nu.Valid {
			return ids, errors.New("state: file UUID not valid")
		}

		ids = append(ids, nu.UUID)
	}

	err = rows.Err()

	return
}
