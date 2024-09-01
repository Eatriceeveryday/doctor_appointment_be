package database

import (
	"BackendTugasAkhir/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase(config config.Config) error {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	return nil
}
