package cfg

import "fmt"

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func GetDBConfig() DBConfig {
	return DBConfig{
		Host:     "localhost",
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

//   const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "PG_USER"
// 	password = "PG_PASS"
// 	dbname   = "PG_DATABASE"
// )
