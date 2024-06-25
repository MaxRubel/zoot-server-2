package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func IncrementWsCount() error {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		return fmt.Errorf("error opening sqlite database: %w", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		UPDATE ws_count
		SET count = count + 1
		WHERE id = 1
	`)

	if err != nil {
		return fmt.Errorf("error updating ws_count: %w", err)
	}
	return nil
}
