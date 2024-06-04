package main

import (
	"fmt"
	"google_auth/handlers"
	"net/http"
)

func main() {
	handlers.LoadEnvs()
	// serve static files (CSS, JS, etc.)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// define the routes for the pages
	http.HandleFunc("/register", handlers.RegisterPageHandler)
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginPageHandler)
	http.HandleFunc("/google", handlers.GoogleHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/user/show", handlers.UserShowPageHandler)
	http.HandleFunc("/user/edit", handlers.UserEditPageHandler)

	fmt.Println("Server listening on 8080")
	http.ListenAndServe(":8080", nil)
}
