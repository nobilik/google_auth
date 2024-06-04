package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
)

type PageData struct {
	Title      string
	User       User
	GoogleLink string
}

type User struct {
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	IsGoogler bool   `json:"is_googler"`
}

const (
	GoogleOAuthURL = "https://accounts.google.com/o/oauth2/v2/auth"
)

var (
	APIHost        string
	Host           string
	GoogleClientID string
	GoogleScope    = "openid https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile"
	client         = http.Client{
		Timeout: 30 * time.Second,
	}
)

func LoadEnvs() {
	GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	APIHost = os.Getenv("API_HOST")
	Host = os.Getenv("HOST")
}

func googleLink() string {
	return fmt.Sprintf("%s?scope=%s&response_type=code&include_granted_scopes=true&prompt=consent&access_type=offline&redirect_uri=%s/google&client_id=%s", GoogleOAuthURL, GoogleScope, Host, GoogleClientID)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	tmplFile := fmt.Sprintf("templates/%s.html", tmpl)
	t, err := template.ParseFiles("templates/layout.html", tmplFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func proccessUserGet(w http.ResponseWriter, r *http.Request, template, title string) {
	req, _ := http.NewRequest(http.MethodGet, APIHost+"/user", nil)
	req.Header.Add("Cookie", r.Header.Get("Cookie"))
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		http.Error(w, string(body), res.StatusCode)
		return
	}

	var user User
	json.Unmarshal(body, &user)

	w.WriteHeader(http.StatusOK)
	renderTemplate(w, template, PageData{Title: title, User: user})
}
