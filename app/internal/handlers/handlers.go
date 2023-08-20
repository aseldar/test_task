package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/aseldar/test_task/app/db"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	Created   time.Time `json:"created"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	db, err := db.ConnectDB()
	if err != nil {
		log.Printf("Error initializing DB: %v", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	user.ID = uuid.New()
	user.Created = time.Now()

	result, err := db.Exec(`
	  INSERT INTO users (id, firstname, lastname, email, age, created)
	  VALUES ($1, $2, $3, $4, $5, $6)`,
		user.ID, user.Firstname, user.Lastname, user.Email, user.Age, user.Created)

	if err != nil {
		log.Printf("Error inserting user: %v", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	if rowsAffected != 1 {
		log.Printf("Error inserting user: no rows affected")
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	db, err := db.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := chi.URLParam(r, "id")

	var user User

	err = db.QueryRow("SELECT id, firstname, lastname, email, age, created FROM users WHERE id = $1", id).Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Age, &user.Created)

	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	db, err := db.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := chi.URLParam(r, "id")

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec(
		"UPDATE users SET firstname = $1, lastname = $2, email = $3, age = $4 WHERE id = $5",
		user.Firstname, user.Lastname, user.Email, user.Age, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
