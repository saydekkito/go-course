package migrations

import (
	"log"
)

func CreateRolesTable() {
	dropRolesTableQuery := `
		DROP TABLE IF EXISTS roles;
	`
	dropRolesIndexQuery := `
		DROP INDEX IF EXISTS roles_title_unique;
	`
	createRolesTableQuery := `
		CREATE TABLE roles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			CONSTRAINT roles_title_not_empty_or_whitespaces CHECK (LENGTH(TRIM(title)) > 0)
		);
	`
	createRolesIndexQuery := `
		CREATE UNIQUE INDEX roles_title_unique 
		ON roles(LOWER(title));
	`

	var err error
	if _, err = db.Exec(dropRolesIndexQuery); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(dropRolesTableQuery); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(createRolesTableQuery); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(createRolesIndexQuery); err != nil {
		log.Fatal(err)
	}
}
