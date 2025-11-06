package bird_species

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/saydekkito/go-course/database"
	"github.com/saydekkito/go-course/models"
)

func UpdateBirdSpecies(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

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

	birdSpeciesID, _ := strconv.Atoi(id)
	_, err = database.DB.Exec("UPDATE bird_species SET title = ?, description = ? WHERE id = ?", bs.Title, bs.Description, birdSpeciesID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs.ID = birdSpeciesID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bs)
}
