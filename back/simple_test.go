package main

import (
	"google_auth/api"
	"google_auth/auth"
	"google_auth/database"
	"google_auth/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// func startup() {

// }

func TestLoginHandler_ValidCredentials(t *testing.T) {
	database.Connect()
	api.Server()
	defer database.DB.Close()

	user := &models.User{
		Email:    "test@example.com",
		Password: "password123",
	}
	err := database.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("email", "test@example.com")
	q.Add("password", "password123")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(auth.LoginHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, rr.Code)
	}
}

func TestRegisterHandler(t *testing.T) {
	database.Connect()
	api.Server()
	defer database.DB.Close()

	q := url.Values{}
	q.Add("email", "test@example.com")
	q.Add("password", "password123")

	req, err := http.NewRequest("POST", "/register", strings.NewReader(q.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(auth.RegisterHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusOK, rr.Code)
	}
}
