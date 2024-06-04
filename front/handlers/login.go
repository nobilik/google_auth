package handlers

import (
	"io"
	"net/http"
	"strings"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		renderTemplate(w, "login", PageData{Title: "Login", GoogleLink: googleLink()})
	case http.MethodPost:
		processLogin(w, r)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

}

func processLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	req, _ := http.NewRequest(http.MethodPost, APIHost+"/login", strings.NewReader(r.Form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		http.Error(w, string(body), res.StatusCode)
		return
	}
	w.Header().Add("Set-Cookie", res.Header.Get("Set-Cookie"))
	w.WriteHeader(http.StatusOK)
}
