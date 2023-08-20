package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aseldar/test_task/app/internal/handlers"
)

var testUser = `{
	"firstname": "Test",
	"lastname": "User",
	"email": "test.user@example.com",
	"age": 30
	}`

func TestCreateUser(t *testing.T) {

	// подготовка данных

	rec := httptest.NewRecorder()

	testUserBytes := []byte(testUser)

	req := httptest.NewRequest(
		"POST",
		"/users",
		bytes.NewBuffer(testUserBytes),
	)

	// вызов хендлера

	handlers.CreateUser(rec, req)

	// проверки

	if status := rec.Code; status != http.StatusCreated {
		t.Errorf("Got wrong status code: %v", status)
	}

	// другие проверки ответа
}
