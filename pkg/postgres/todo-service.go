package postgres

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type ToDoService struct {
	DbUserName string
	DbPassword string
	DbURL      string
	DbName     string
}

func (t *ToDoService) Create(text string, isDone bool) {
	panic("oops!")
}

func (t *ToDoService) Initialise() error {
	dbConnectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", t.DbUserName, t.DbPassword, t.DbURL, t.DbName)
	db, err := sql.Open("pgx", dbConnectionString)
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://../../pkg/postgres/migrations", t.DbName, driver)
	if err != nil {
		return err
	}
	err = m.Steps(2)
	if err != nil {
		return err
	}
	return nil
}
