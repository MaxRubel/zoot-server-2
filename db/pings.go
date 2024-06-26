package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func AddPing() error {

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}
	defer db.Close()
	query := `UPDATE pings SET number = number + 1 WHERE Id = 1;`
	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("error incrementing ping in db: %w", err)
	}

	return nil
}
