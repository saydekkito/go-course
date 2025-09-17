package main

import (
	"log"
	"net/http"

	"github.com/saydekkito/go-course/database"
	"github.com/saydekkito/go-course/routes"
)

func main() {
	database.Connect()
	database.InitDB()
	database.SeedDB()

	r := routes.SetupRouter()
	log.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8080", r)
}
