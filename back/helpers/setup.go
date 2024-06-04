package helpers

import (
	"os"
)

var (
	GoogleClientID     string
	GoogleClientSecret string
)

func LoadEnv() {
	GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
}
