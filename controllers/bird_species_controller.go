package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

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

func DeleteBirdSpecies(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	birdSpeciesID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Некорректный ID вида птиц", http.StatusBadRequest)
	}

	result, err := database.DB.Exec("DELETE FROM bird_species WHERE id = ?", birdSpeciesID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Вид птицы не найден"})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"message": "Вид птицы успешно удалён"})
	}
}
