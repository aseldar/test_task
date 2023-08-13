package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	Created   time.Time `json:"created"`
}

var db *sql.DB

func main() {
	// Подключение к базе данных PostgreSQL
	db, err := sql.Open("postgres", "postgres://username:password@postgres:5432/dbname?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Подключение к миграциям
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	// Создание маршрутов для API
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")

	// Запуск сервера
	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = uuid.New()
	user.Created = time.Now()

	_, err := db.Exec(
		"INSERT INTO users (id, firstname, lastname, email, age, created) VALUES ($1, $2, $3, $4, $5, $6)",
		user.ID, user.Firstname, user.Lastname, user.Email, user.Age, user.Created,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var user User
	row := db.QueryRow("SELECT id, firstname, lastname, email, age, created FROM users WHERE id = $1", id)
	if err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Age, &user.Created); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec(
		"UPDATE users SET firstname = $1, lastname = $2, email = $3, age = $4 WHERE id = $5",
		user.Firstname, user.Lastname, user.Email, user.Age, id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
