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

// @title           Swagger Example API
// @version         1.0
// @description     This is a test task effetive-mobile.
// @termsOfService

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  sgk1988@yandex.ru

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	err := config.Init()
	if err != nil {
		fmt.Print(err)
		return
	}
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

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(dbType); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}
