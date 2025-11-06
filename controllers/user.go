package controllers

import (
	"database/sql"
	"log"

	"github.com/saydekkito/go-course/database"
	"github.com/saydekkito/go-course/models"
)

func FindUser(username string) *models.User {
	var user models.User
	var role models.Role

	query := `
		SELECT u.id, u.username, u.password, u.role_id, r.id, r.title
		FROM users u
		JOIN roles r ON u.role_id = r.id
		WHERE LOWER(u.username) = LOWER(?)
	`

	err := database.DB.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.RoleID,
		&role.ID,
		&role.Title,
	)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("Ошибка при поиске пользователя:", err)
		}
		return nil
	}

	user.Role = &role
	return &user
}
