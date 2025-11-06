package migrations

import "log"

func CreateBirdSpeciesTable() {
	dropBirdSpeciesTableQuery := `
		DROP TABLE IF EXISTS bird_species;
	`
	dropBirdSpeciesIndexQuery := `
		DROP INDEX IF EXISTS bird_species_title_unique;
	`
	createBirdSpeciesTableQuery := `
		CREATE TABLE bird_species (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			CONSTRAINT bird_species_title_not_empty_or_whitespaces CHECK (LENGTH(TRIM(title)) > 0)
		);
	`
	createBirdSpeciesIndexQuery := `
		CREATE UNIQUE INDEX bird_species_title_unique 
		ON bird_species(LOWER(title));
	`

	var err error
	if _, err = db.Exec(dropBirdSpeciesIndexQuery); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(dropBirdSpeciesTableQuery); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(createBirdSpeciesTableQuery); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(createBirdSpeciesIndexQuery); err != nil {
		log.Fatal(err)
	}
}
