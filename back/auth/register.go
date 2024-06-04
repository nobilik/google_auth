package auth

import (
	"google_auth/database"
	"google_auth/models"
	"net/http"
)

// register with email & password
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if IsAuthenticated(r) {
		http.Error(w, "Already logged in", http.StatusForbidden)
		return
	}

	var user models.User

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")

	err = database.CreateUser(&user)
	if err != nil {
		if database.IsDuplicateEntryError(err) {
			http.Error(w, "Already exists", http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Log user in after registration
	session, err := store.Get(r, "test_session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["authenticated"] = true
	session.Values["user_id"] = user.ID
	session.Save(r, w)
	w.WriteHeader(http.StatusCreated)
}
