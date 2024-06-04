package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	store = sessions.NewCookieStore([]byte("test_secret_key"))
)

func IsAuthenticated(r *http.Request) bool {
	session, err := store.Get(r, "test_session")
	if err != nil {
		return false
	}
	auth, ok := session.Values["authenticated"].(bool)
	return ok && auth
}

func GetUserID(r *http.Request) uint {
	session, err := store.Get(r, "test_session")
	if err != nil {
		return 0
	}
	id, ok := session.Values["user_id"].(uint)
	if !ok {
		return 0
	}
	return id
}
