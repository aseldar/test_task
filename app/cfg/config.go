package cfg

import (
	"fmt"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// var POSTGRES_USER string = os.Getenv("POSTGRES_USER")
// var POSTGRES_PASSWORD string = os.Getenv("POSTGRES_PASSWORD")
// var POSTGRES_DB string = os.Getenv("POSTGRES_DB")

func GetDBConfig() DBConfig {
	return DBConfig{
		Host:     "postgres",
		Port:     5432,
		User:     "PG_USER",
		Password: "PG_PASS",
		DBName:   "PG_DATABASE",
	}
}

func (c *DBConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)
}
