package controllers

import (
	"database/sql"
	"log"

	"github.com/saydekkito/go-course/database"
	"github.com/saydekkito/go-course/models"
)

func FindRole(roleTitle string) *models.Role {
	var role models.Role

	query := `
		SELECT id, string
		FROM roles
		WHERE title = ?;
	`

	err := database.DB.QueryRow(query, roleTitle).Scan(
		&role.ID,
		&role.Title,
	)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("Ошибка при поиске пользователя:", err)
		}
		return nil
	}

	return &role
}
