package migrations

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreatedb, downCreatedb)
}

func upCreatedb(ctx context.Context, tx *sql.Tx) error {
	db := os.Getenv("DB")
	owner := os.Getenv("OWNER")
	creatDB := fmt.Sprintf(
		`CREATE DATABASE %v
		WITH
		OWNER = %v
		ENCODING = 'UTF8'
		CONNECTION LIMIT = -1;`,
		db,
		owner,
	)
	_, err := tx.ExecContext(ctx, creatDB)
	if err != nil {
		return err
	}

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	addUser := fmt.Sprintf(
		"CREATE USER %v WITH PASSWORD '%v';",
		user,
		password,
	)

	_, err = tx.ExecContext(ctx, addUser)
	if err != nil {
		return err
	}

	return nil

}

func downCreatedb(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
