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
    - bird_species: контроллеры для видов птиц;
    - auth.go: контроллер для авторизации;
    - role.go: контроллер для ролей;
    - user.go: контроллер для пользователей.
- database:
    - migrations: миграции;
    - connect.go: реализация подключения к БД;
    - seeder.go: наполнение БД.
- middleware:
    - jwt.go: middleware для JWT-токенов.
- models:
    - bird_species.go: модель видов птиц;
    - role.go: модель ролей;
    - user.go: модель пользователей.
- queries: тестовые запросы.
- routes:
    - routes.go: endpoints API.
- utils:
    - env.go: util для работы с переменными окружения.


---
Инструкция к использованию проекта:
1. git clone https://github.com/saydekkito/go-course.git
2. cd go-course
3. go mod tidy
4. touch .env
5. Добавить в файл .env переменные PORT и JWT_SECRET.
6. go run main.go
7. Тестовые запросы лежат в папке queries.

---