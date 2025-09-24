API «Виды птиц»
---
Описание:
REST API проект, предоставляющий взаимодействие с базой данных видов птиц.

---
Таблица «bird_species» (виды птиц):
| Атрибут       | Тип           | Описание                        |
| ------------- | ------------- | ------------------------------- |
| id            | INTEGER (PK)  | Первичный ключ                  |
| title         | TEXT          | Название вида (уникальное)      |
| description   | TEXT          | Описание вида (может быть null) |

---
Структура проекта:
- controllers:
    - bird_species_controller.go: контроллер для видов птиц.
- database:
    - database/database.go: реализация подключения к БД и создания таблицы;
    - database/seed.go: наполнение БД.
- models:
    - bird_species.go: модель видов птиц.
- queries: тестовые запросы.
- routes:
    - routes.go: endpoints API.


---
Инструкция к использованию проекта:
1. git clone https://github.com/saydekkito/go-course.git
2. cd go-course
3. go mod tidy
4. go run main.go
5. Тестовые запросы лежат в папке queries.

---