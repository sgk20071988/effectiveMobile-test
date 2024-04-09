package context

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func GetDB() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	return nil, fmt.Errorf("no sql client")
}

func SetDB(_db *sql.DB) {
	db = _db
}
