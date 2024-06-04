package pages

import (
	"google_auth/auth"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	// // Authentication check
	if !auth.IsAuthenticated(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	switch r.Method {
	case http.MethodGet:
		getUserHandler(w, r)
	case http.MethodPut:
		updateUserHandler(w, r)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}
