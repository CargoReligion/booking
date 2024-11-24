package db

import (
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DbConnectionInfo struct {
	Host         string
	Port         string
	Username     string
	DatabaseName string
}

func GetDbConnectionInfo() *DbConnectionInfo {
	return &DbConnectionInfo{
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		Username:     os.Getenv("DB_USER"),
		DatabaseName: os.Getenv("DB_NAME"),
	}
}

func GetDbConnection(attempts int) (*sqlx.DB, error) {
	connInfo := GetDbConnectionInfo()

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		connInfo.Host,
		connInfo.Port,
		connInfo.Username,
		os.Getenv("DB_PASSWORD"),
		connInfo.DatabaseName,
	)

	var db *sqlx.DB
	var err error

	for i := 0; i < attempts; i++ {
		wait(250 * i)

		db, err = sqlx.Connect("postgres", psqlInfo)
		if err != nil {
			continue
		}

		err = db.Ping()
		if err == nil {
			break
		}
	}

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(10 * time.Minute)
	return db, nil
}

func wait(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
