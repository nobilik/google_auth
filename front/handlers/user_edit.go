package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func UserEditPageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		proccessUserGet(w, r, "user_edit", "Edit User")
	case http.MethodPut:
		updateUser(w, r)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	rBody, _ := io.ReadAll(r.Body)
	req, _ := http.NewRequest(http.MethodPut, APIHost+"/user", bytes.NewBuffer(rBody))
	req.Header.Add("Cookie", r.Header.Get("Cookie"))
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Print(res.StatusCode)
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized {
			http.Redirect(w, r, "/login", http.StatusSeeOther)

			return
		}
		http.Error(w, string(body), res.StatusCode)
		return
	}
	w.WriteHeader(http.StatusOK)
}
