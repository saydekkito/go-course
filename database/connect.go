package database

import (
	"database/sql"
	"log"

	"github.com/saydekkito/go-course/database/migrations"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect() {
	var err error
	if DB, err = sql.Open("sqlite", "./birds.DB"); err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Подключение к БД: успех")
}

func InitDB() {
	migrations.SetDB(DB)
	migrations.RunMigrations()
}
