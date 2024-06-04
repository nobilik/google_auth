package handlers

import (
	"io"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
	req, _ := http.NewRequest(http.MethodDelete, APIHost+"/logout", nil)
	req.Header.Add("Cookie", r.Header.Get("Cookie"))
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
	w.WriteHeader(http.StatusOK)
}
