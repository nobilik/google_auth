package handlers

import (
	"io"
	"net/http"
	"strings"
)

func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		renderTemplate(w, "register", PageData{Title: "Register", GoogleLink: googleLink()})
	case http.MethodPost:
		processRegister(w, r)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func processRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	req, _ := http.NewRequest(http.MethodPost, APIHost+"/register", strings.NewReader(r.Form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusCreated {
		http.Error(w, string(body), res.StatusCode)
		return
	}
	w.Header().Add("Set-Cookie", res.Header.Get("Set-Cookie"))
	w.WriteHeader(http.StatusCreated)
}
