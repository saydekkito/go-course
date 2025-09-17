package routes

import (
	"github.com/gorilla/mux"
	"github.com/saydekkito/go-course/controllers"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/bird_species", controllers.GetAllBirdSpecies).Methods("GET")
	r.HandleFunc("/bird_species/{id}", controllers.GetBirdSpecies).Methods("GET")
	r.HandleFunc("/bird_species", controllers.CreateBirdSpecies).Methods("POST")
	r.HandleFunc("/bird_species/{id}", controllers.UpdateBirdSpecies).Methods("PUT")
	r.HandleFunc("/bird_species/{id}", controllers.DeleteBirdSpecies).Methods("DELETE")

	return r
}
