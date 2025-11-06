package bird_species

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/saydekkito/go-course/database"
	"github.com/saydekkito/go-course/models"
)

func CreateBirdSpecies(w http.ResponseWriter, r *http.Request) {
	var bs models.BirdSpecies
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &bs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := database.DB.Exec("INSERT INTO bird_species(title, description) VALUES(?, ?)", bs.Title, bs.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	bs.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bs)
}
