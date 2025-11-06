package main

import (
	"log"
	"net/http"

	"github.com/saydekkito/go-course/database"
	"github.com/saydekkito/go-course/routes"
	"github.com/saydekkito/go-course/utils"
)

func main() {
	database.Connect()
	database.InitDB()
	database.SeedDB()

	port := utils.GetEnv("PORT", "8080")

	r := routes.SetupRouter()
	log.Println("Запуск сервера: сервер запущен на", port)
	http.ListenAndServe(":"+port, r)
}
