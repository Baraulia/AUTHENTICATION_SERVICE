package database

import (
	"database/sql"
	"fmt"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
)

type PostgresDB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
	logger   logging.Logger
}

func NewPostgresDB(database PostgresDB) (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			database.Host, database.Port, database.Username, database.DBName, database.Password, database.SSLMode))
	if err != nil {
		database.logger.Panicf("Database open error:%s", err)
		return nil, fmt.Errorf("error connecting to database:%s", err)
	}
	err = db.Ping()
	if err != nil {
		database.logger.Errorf("DB ping error:%s", err)
		return nil, err
	}
	_, err = db.Exec(USER_SCHEMA)
	if err != nil {
		database.logger.Errorf("Error executing initial migration into users:%s", err)
		return nil, fmt.Errorf("error executing initial migration into users:%s", err)
	}
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