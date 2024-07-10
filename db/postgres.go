package db

import (
	"database/sql"
	"fmt"
	"forwardchaining/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type Postgres struct{}

func Connect(creds *model.Credential) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func SQLExecute(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username VARCHAR(255), password VARCHAR(255))")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS sessions (id SERIAL PRIMARY KEY, token VARCHAR(255), username VARCHAR(255), expiry TIMESTAMP DEFAULT NULL)")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS diseases (id SERIAL PRIMARY KEY, code VARCHAR(255) UNIQUE, name VARCHAR(255) UNIQUE)")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS symptoms (id SERIAL PRIMARY KEY, code VARCHAR(255) UNIQUE, name VARCHAR(255))")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS rules (id SERIAL PRIMARY KEY, code_diseases VARCHAR(255) REFERENCES diseases(code), code_symptoms VARCHAR(255) REFERENCES symptoms(code), md DECIMAL(10,2), mb DECIMAL(10,2))")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS questions (id SERIAL PRIMARY KEY, code VARCHAR(255), question VARCHAR(255))")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS diagnoses (id SERIAL PRIMARY KEY, name VARCHAR(255) REFERENCES diseases(name), nilai DECIMAL(10, 2), description VARCHAR(255))")
	if err != nil {
		return err
	}

	return nil
}

func Reset(db *sql.DB, table string) error {
	_, err := db.Exec("TRUNCATE " + table)
	if err != nil {
		return err
	}

	_, err = db.Exec("ALTER SEQUENCE " + table + "_id_seq RESTART WITH 1")
	if err != nil {
		return err
	}

	return nil
}
