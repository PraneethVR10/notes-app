package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func DBCongif() *DB {
	return &DB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

}

func ConnectDB() (*sql.DB, error) {
	config := DBCongif()

	if config.Host == "" || config.Port == "" || config.User == "" ||
		config.Password == "" || config.DBName == "" {
		return nil, fmt.Errorf("missing required database environment variables")

	} else {
		connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Host,
			config.Port,
			config.User,
			config.Password,
			config.DBName,
		)
		db, err := sql.Open("postgres", connect)

		if err != nil {
			return nil, err
		}

		err = db.Ping()
		if err != nil {
			return nil, err
		}

		log.Printf("Successfully connected to database!\n")
		return db, nil
	}

}
