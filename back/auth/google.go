package auth

import (
	"fmt"
	"google_auth/database"
	"google_auth/external"
	"google_auth/models"
	"net/http"
)

// we use it for login and register
func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	redirectURI := r.URL.Query().Get("redirect_uri")

	tokenData, err := external.GetGoogleOauthToken(redirectURI, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data, err := external.GetGoogleUserInfo(tokenData["id_token"].(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user, err := database.GetUserByGoogleUID(data["sub"].(string))
	// it's registered user - login
	if err == nil {
		session, err := store.Get(r, "test_session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["authenticated"] = true
		session.Values["user_id"] = user.ID
		session.Save(r, w)

		w.WriteHeader(http.StatusOK)
		return
	}

	// we try to create user with google email
	user = &models.User{
		Email:    data["email"].(string),
		FullName: fmt.Sprintf("%s %s", data["given_name"].(string), data["family_name"].(string)),
	}
	err = database.CreateUser(user)

	if err != nil {
		// case user was already created with email|pass
		if database.IsDuplicateEntryError(err) {
			http.Error(w, "Already exists", http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	profile := &models.GoogleProfile{
		UserID: user.ID,
		UID:    data["sub"].(string),
	}

	database.CreateGoogleProfile(profile)

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
