package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aseldar/test_task/app/db"
	"github.com/aseldar/test_task/app/internal/handlers"

	"github.com/go-chi/chi"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	db, err := db.InitializeDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// // Создание маршрутов для API
	r := chi.NewRouter()

	r.Get("/", healthCheck)
	r.Get("/users/{id}", handlers.GetUser)
	r.Post("/users", handlers.CreateUser)
	r.Put("/users/{id}", handlers.UpdateUser)

	// Запуск сервера
	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
	})
}
