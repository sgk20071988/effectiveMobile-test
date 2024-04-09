package main

import (
	"database/sql"
	"effectiveMobileTest/internal/config"
	"embed"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	err := config.Init()
	if err != nil {
		fmt.Print(err)
		return
	}
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dbType := os.Getenv("DB_TYPE")
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSource := fmt.Sprintf(
		"%v://%v:%v@%v:%v/%v",
		dbType,
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
	fmt.Println(dbSource)

	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		panic(err)
	}
	/*
		pingErr := db.Ping()
		if pingErr != nil {
			panic(pingErr)
		}
		fmt.Println("Connected!")
	*/
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(dbType); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}
