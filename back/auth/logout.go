package auth

import "net/http"

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "test_session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Revoke user authentication
	session.Values["authenticated"] = false
	session.Values["user_id"] = 0
	session.Save(r, w)

	w.WriteHeader(http.StatusOK)
}
