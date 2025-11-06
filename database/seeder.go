package database

import (
	"log"

	"github.com/saydekkito/go-course/models"
)

func SeedDB() {
	bird_species := []models.BirdSpecies{
		{Title: "Украшенный чибис", Description: "Редкая птица из семейства ржанковых, отличается хохолком и контрастным оперением."},
		{Title: "Чибис", Description: "Обычная луговая птица Евразии, известная своим «кувыркающимся» полетом и звуками."},
		{Title: "Золотистая ржанка", Description: "Мелкая перелётная птица с золотистыми пятнами на спине, гнездится в тундре."},
		{Title: "Тулес", Description: "Европейский кулик, часто встречается на болотах и поймах рек."},
		{Title: "Азиатский бекас", Description: "Перелётный кулик, распространён в Азии, любит заболоченные луга и берега водоёмов."},
		{Title: "Бекас", Description: "Характерен длинным прямым клювом и зигзагообразным полётом, часто слышен весной."},
		{Title: "Азиатский бекасовидный веретенник", Description: "Редкий вид, напоминающий веретенника, но с чертами бекаса."},
		{Title: "Короткоклювый бекасовидный веретенник", Description: "Отличается более коротким клювом, встречается на болотах Восточной Азии."},
		{Title: "Альпийская галка", Description: "Горная птица, обитающая в Альпах и Гималаях, известна акробатическим полётом."},
		{Title: "Клушица", Description: "Крупная ворона с красным клювом и ногами, встречается в горах Евразии."},
	}

	for _, bs := range bird_species {
		_, err := DB.Exec("INSERT INTO bird_species(title, description) VALUES(?, ?)", bs.Title, bs.Description)
		if err != nil {
			log.Println("Наполнение базы: ошибка: ", err)
		}
	}

	roles := []models.Role{
		{Title: "admin"},
		{Title: "user"},
	}

	for _, role := range roles {
		_, err := DB.Exec("INSERT INTO roles(title) VALUES(?)", role.Title)
		if err != nil {
			log.Println("Наполнение базы: ошибка: ", err)
		}
	}

	var adminRoleID, userRoleID int
	err := DB.QueryRow("SELECT id FROM roles WHERE title = ?", "admin").Scan(&adminRoleID)
	if err != nil {
		log.Fatal("Ошибка при получении ID роли admin:", err)
	}

	err = DB.QueryRow("SELECT id FROM roles WHERE title = ?", "user").Scan(&userRoleID)
	if err != nil {
		log.Fatal("Ошибка при получении ID роли user:", err)
	}

	users := []models.User{
		{Username: "admin", Password: "admin", RoleID: adminRoleID},
		{Username: "user", Password: "user", RoleID: userRoleID},
	}

	for _, user := range users {
		_, err := DB.Exec("INSERT INTO users(username, password, role_id) VALUES(?, ?, ?)", user.Username, user.Password, user.RoleID)
		if err != nil {
			log.Println("Наполнение базы: ошибка: ", err)
		}
	}

	log.Println("Наполнение базы: успех")
}
