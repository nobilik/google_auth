package auth

import (
	"google_auth/database"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	email, password := r.FormValue("email"), r.FormValue("password")

	user, err := database.GetUser("email", email)
	if err != nil {
		if database.IsNotFoundError(err) {
			http.Error(w, "username entered does not exist", http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if password != user.Password {
		http.Error(w, "password is incorrect", http.StatusUnauthorized)
		return
	}

	// Create a new session
	session, err := store.Get(r, "test_session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["authenticated"] = true
	session.Values["user_id"] = user.ID
	session.Save(r, w)

	w.WriteHeader(http.StatusOK)
}
