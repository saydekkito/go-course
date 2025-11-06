package bird_species

// TODO: добавить валидацию данных для всех контроллеров

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/saydekkito/go-course/database"
	"github.com/saydekkito/go-course/models"
)

func GetAllBirdSpecies(w http.ResponseWriter, r *http.Request) {
	titleFilter := r.URL.Query().Get("title")
	descFilter := r.URL.Query().Get("description")
	sortOrder := strings.ToUpper(r.URL.Query().Get("sort"))
	if sortOrder != "DESC" {
		sortOrder = "ASC"
	}

	query := `SELECT id, title, description FROM bird_species`
	var conditions []string
	var args []interface{}

	if titleFilter != "" {
		conditions = append(conditions, "title LIKE ?")
		args = append(args, "%"+titleFilter+"%")
	}
	if descFilter != "" {
		conditions = append(conditions, "description LIKE ?")
		args = append(args, "%"+descFilter+"%")
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY title " + sortOrder

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		http.Error(w, "Ошибка при запросе к БД", http.StatusInternalServerError)
		log.Println("Ошибка при запросе к БД:", err)
		return
	}
	defer rows.Close()

	var birdSpecies []models.BirdSpecies
	for rows.Next() {
		var bs models.BirdSpecies
		if err := rows.Scan(&bs.ID, &bs.Title, &bs.Description); err != nil {
			log.Println("Ошибка сканирования строки:", err)
			continue
		}
		birdSpecies = append(birdSpecies, bs)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(birdSpecies)
}

func GetBirdSpecies(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var bs models.BirdSpecies
	database.DB.QueryRow("SELECT id, title, description FROM bird_species WHERE id = ?", id).Scan(&bs.ID, &bs.Title, &bs.Description)
	json.NewEncoder(w).Encode(bs)
}
