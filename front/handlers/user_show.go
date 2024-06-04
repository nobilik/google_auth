package handlers

import (
	"net/http"
)

func UserShowPageHandler(w http.ResponseWriter, r *http.Request) {
	proccessUserGet(w, r, "user_show", "Show User")
}
