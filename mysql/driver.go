package mysql

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Database *sqlx.DB

func Init() {
	dataSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sqlx.Connect("mysql", dataSource)

	if err != nil {
		panic(err.Error())
	}

	Database = db
}

func Close() {
	Database.Close()
}
