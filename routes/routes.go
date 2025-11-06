package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saydekkito/go-course/controllers"
	"github.com/saydekkito/go-course/controllers/bird_species"
	"github.com/saydekkito/go-course/middleware"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/bird_species", bird_species.GetAllBirdSpecies).Methods("GET")
	r.HandleFunc("/bird_species/{id}", bird_species.GetBirdSpecies).Methods("GET")

	r.HandleFunc("/bird_species", middleware.JWTAuth(http.HandlerFunc(bird_species.CreateBirdSpecies))).Methods("POST")
	r.HandleFunc("/bird_species/{id}", middleware.JWTAuth(http.HandlerFunc(bird_species.UpdateBirdSpecies))).Methods("PUT")
	r.HandleFunc("/bird_species/{id}", middleware.JWTAuth(http.HandlerFunc(bird_species.DeleteBirdSpecies))).Methods("DELETE")

	return r
}
