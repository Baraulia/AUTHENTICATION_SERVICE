package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"

	"github.com/spf13/viper"
)

var DB *sql.DB

func SetupDB() (*sql.DB, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	// Find and read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	userName := viper.GetString("APP_DB_USERNAME")
	password := viper.GetString("APP_DB_PASSWORD")
	hostName := viper.GetString("APP_DB_HOST")
	dbName := viper.GetString("APP_DB_NAME")
	dbPort := viper.GetString("APP_DB_PORT")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",userName, password, hostName, dbPort, dbName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(USER_SCHEMA)

	return db, nil
}

const USER_SCHEMA = `
	CREATE TABLE IF NOT EXISTS users (
		id serial not null primary key ,
		email varchar(225) NOT NULL UNIQUE,
		password varchar(225) NOT NULL,
	    activated  boolean NOT NULL default false,
		created_at timestamp NOT NULL,
		updated_at timestamp NOT NULL
	);
`