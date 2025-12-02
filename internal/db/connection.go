package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
    dbUsername := os.Getenv("DB_USERNAME")
    dbPassword := os.Getenv("DB_PASSWORD")

    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
    log.Print(connectionString)
    return sql.Open("mysql",
    	fmt.Sprintf("%s?parseTime=true", connectionString),
	)
}

