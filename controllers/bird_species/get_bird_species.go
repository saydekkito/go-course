package bird_species

// TODO: добавить валидацию данных для всех контроллеров
// TODO: допилить документацию
// TODO: добавить фильтры, поиск по полям для get запросов

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saydekkito/go-course/database"
	"github.com/saydekkito/go-course/models"
)

func GetAllBirdSpecies(w http.ResponseWriter, r *http.Request) {
	rows, _ := database.DB.Query("SELECT id, title, description FROM bird_species;")
	defer rows.Close()

	var bird_species []models.BirdSpecies
	for rows.Next() {
		var bs models.BirdSpecies
		rows.Scan(&bs.ID, &bs.Title, &bs.Description)
		bird_species = append(bird_species, bs)
	}
	json.NewEncoder(w).Encode(bird_species)
}

func GetBirdSpecies(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var bs models.BirdSpecies
	database.DB.QueryRow("SELECT id, title, description FROM bird_species WHERE id = ?", id).Scan(&bs.ID, &bs.Title, &bs.Description)
	json.NewEncoder(w).Encode(bs)
}
