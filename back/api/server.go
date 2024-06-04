package api

import (
	"fmt"

	"google_auth/auth"
	"google_auth/pages"
	"net/http"
)

func Server() {
	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/logout", auth.LogoutHandler)
	http.HandleFunc("/register", auth.RegisterHandler)
	http.HandleFunc("/google", auth.GoogleLogin)
	http.HandleFunc("/user", pages.UserHandler)

	fmt.Println("Server listening on 3000")
	http.ListenAndServe(":3000", nil)

}
