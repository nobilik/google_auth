package pages

import (
	"encoding/json"
	"google_auth/auth"
	"google_auth/database"
	"google_auth/helpers"
	"google_auth/models"
	"net/http"
)

// for user's handlers we take ID from session
func getUserHandler(w http.ResponseWriter, r *http.Request) {

	userID := auth.GetUserID(r)
	if userID == 0 {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	user, err := database.GetUser("id", userID)
	if err != nil {
		if database.IsNotFoundError(err) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.ResponseJSON(w, user)
}

// PUT /user
func updateUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
	user.ID = auth.GetUserID(r)
	if user.ID == 0 {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = database.UpdateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.ResponseJSON(w, user)
}
