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
	user := os.Getenv("USER")
	creatDB := fmt.Sprintf(
		`CREATE DATABASE %v
		WITH
		OWNER = %v
		ENCODING = 'UTF8'
		CONNECTION LIMIT = -1;`,
		db,
		user,
	)
	_, err := tx.ExecContext(ctx, creatDB)
	return err
}

func downCreatedb(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
