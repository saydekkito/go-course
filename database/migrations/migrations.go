package migrations

import "database/sql"

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func RunMigrations() {
	CreateRolesTable()
	CreateUsersTable()
	CreateBirdSpeciesTable()
}
