package external

import (
	"bytes"
	"encoding/json"
	"errors"
	"google_auth/helpers"
	"io"
	"net/http"
	"net/url"
	"time"
)

var (
	client = http.Client{
		Timeout: time.Second * 30,
	}
)

func GetGoogleOauthToken(redirect_uri, code string) (GoogleOauthTokenRes map[string]interface{}, err error) {
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", helpers.GoogleClientID)
	values.Add("client_secret", helpers.GoogleClientSecret)
	values.Add("redirect_uri", redirect_uri)

	query := values.Encode()

	req, err := http.NewRequest("POST", "https://oauth2.googleapis.com/token", bytes.NewBufferString(query))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve token")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resBody.Bytes(), &GoogleOauthTokenRes); err != nil {
		return nil, err
	}

	return GoogleOauthTokenRes, nil
}

func GetGoogleUserInfo(idToken string) (data map[string]interface{}, err error) {
	values := url.Values{}
	values.Add("id_token", idToken)
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/tokeninfo", nil)
	if err != nil {
		return
	}
	req.URL.RawQuery = values.Encode()

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return
	}
	body, _ := io.ReadAll(res.Body)

	if err = json.Unmarshal(body, &data); err != nil {
		return
	}

	return data, nil
}
