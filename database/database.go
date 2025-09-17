package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect() {
	var err error
	if DB, err = sql.Open("sqlite", "./birds.db"); err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Подключение к БД: успех")
}

func InitDB() {
	dropTableQuery := `
	DROP TABLE IF EXISTS bird_species;
	`
	dropIndexQuery := `
	DROP INDEX IF EXISTS bird_species_title_unique;
	`
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS bird_species (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		CONSTRAINT bird_species_title_not_empty_or_whitespaces CHECK (LENGTH(TRIM(title)) > 0)
	);
	`
	createIndexQuery := `
	CREATE UNIQUE INDEX IF NOT EXISTS bird_species_title_unique 
    ON bird_species(LOWER(title));
	`

	var err error
	if _, err = DB.Exec(dropIndexQuery); err != nil {
		log.Fatal(err)
	}
	if _, err = DB.Exec(dropTableQuery); err != nil {
		log.Fatal(err)
	}
	if _, err = DB.Exec(createTableQuery); err != nil {
		log.Fatal(err)
	}
	if _, err = DB.Exec(createIndexQuery); err != nil {
		log.Fatal(err)
	}
}
