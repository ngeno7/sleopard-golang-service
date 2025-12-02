package db


import (
    "database/sql"
   _ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {

    return sql.Open("mysql",
    	"root:root@tcp(127.0.0.1:3306)/lpd?parseTime=true",
	)
}

