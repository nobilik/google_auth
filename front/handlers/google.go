package handlers

import (
	"io"
	"net/http"
)

func GoogleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
	req, _ := http.NewRequest(http.MethodGet, APIHost+"/google", nil)

	q := r.URL.Query()
	q.Add("redirect_uri", Host+"/google")
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode == http.StatusOK {
		w.Header().Add("Set-Cookie", res.Header.Get("Set-Cookie"))
		r.Header["Cookie"] = []string{res.Header.Get("Set-Cookie")}
		UserShowPageHandler(w, r)
		return
	}

	if res.StatusCode == http.StatusCreated {
		w.Header().Add("Set-Cookie", res.Header.Get("Set-Cookie"))
		r.Header["Cookie"] = []string{res.Header.Get("Set-Cookie")}
		UserEditPageHandler(w, r)
		return
	}

	if res.StatusCode == http.StatusConflict {
		RegisterPageHandler(w, r)
		return
	}

	http.Error(w, string(body), res.StatusCode)
}
