package migrations

import (
	"log"
)

func CreateUsersTable() {
	dropUsersTableQuery := `
		DROP TABLE IF EXISTS users;
	`
	dropUsersIndexQuery := `
		DROP INDEX IF EXISTS users_username_unique;
	`
	createUsersTableQuery := `
		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			role_id INTEGER NOT NULL,
			FOREIGN KEY (role_id) REFERENCES roles(id),
			CONSTRAINT users_username_not_empty_or_whitespaces CHECK (LENGTH(TRIM(username)) > 0),
			CONSTRAINT users_password_not_empty_or_whitespaces CHECK (LENGTH(TRIM(password)) > 0)
		);
	`
	createUsersIndexQuery := `
		CREATE UNIQUE INDEX users_username_unique 
		ON users(LOWER(username));
	`

	var err error
	if _, err = db.Exec(dropUsersIndexQuery); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(dropUsersTableQuery); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(createUsersTableQuery); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(createUsersIndexQuery); err != nil {
		log.Fatal(err)
	}
}
