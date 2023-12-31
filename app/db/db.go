package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aseldar/test_task/app/cfg"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func InitializeDB() (*sql.DB, error) {
	conf := cfg.GetDBConfig()
	connectionStr := conf.GetConnectionString()

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		fmt.Println("error sql.Open")
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
		fmt.Println("error WithInstance")
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/db/migration",
		// "file://C:/Users/asildar.magomedov/test_task/app/db/migration",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
		fmt.Println("error migrate")
		return nil, err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("errror migrate ErrNoChange")
		return nil, err
	}
	return db, nil
}

func ConnectDB() (*sql.DB, error) {
	conf := cfg.GetDBConfig()
	connectionStr := conf.GetConnectionString()

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		fmt.Println("errror sql.Open")
		panic(err)
	}
	return db, nil
}
