package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	localUsername string = "root"
	localPassword string = "password"
	localDatabase string = "db_movie"
)

// HubToMySQL
func MySQL() (*sql.DB, error) {

	// errEnv := godotenv.Load()
	// if errEnv != nil {
	// 	log.Fatalf("Error on loading .env file")
	// }

	var (
		username string = os.Getenv("DB_USERNAME")
		password string = os.Getenv("DB_PASSWORD")
		database string = os.Getenv("DB_DATABASE")
		host     string = os.Getenv("DB_HOST")
	)

	var dsn string

	if host == "" {
		dsn = fmt.Sprintf("%v:%v@/%v", localUsername, localPassword, localDatabase)
	} else {
		dsn = fmt.Sprintf("%v:%v@%v/%v", username, password, host, database)
	}
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
