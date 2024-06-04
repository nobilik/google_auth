package main

import (
	"google_auth/api"
	"google_auth/database"
	"google_auth/helpers"
)

func main() {
	helpers.LoadEnv()
	database.Connect()
	defer database.DB.Close()
	api.Server()
}
