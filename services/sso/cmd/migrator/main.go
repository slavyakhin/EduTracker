package main

import (
	"errors"
	"flag"
	"fmt"

	// Migration lib
	"github.com/golang-migrate/migrate/v4"
	// Migration driver for postgres
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	// Migration file driver
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var migrationsPath, migrationsTable string

	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse()

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("postgres://postgres:mysecretpassword@localhost:5432/postgres?x-migrations-table=%s&sslmode=disable", migrationsTable), // TODO: make correct URI
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}

		panic(err)
	}

	fmt.Println("migrations applied successfully")
}
