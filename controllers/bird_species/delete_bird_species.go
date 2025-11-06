package bird_species

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/saydekkito/go-course/database"
)

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
