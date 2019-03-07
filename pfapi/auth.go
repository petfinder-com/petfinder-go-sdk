package pfapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

//BearerToken structure
type BearerToken struct {
	TokenType   string `json:"token_type"`
	Expiration  int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

//Authenticate method accepting client_id and client_secret
func Authenticate(clientID string, clientSecret string) (string, error) {
	form := url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"grant_type":    {"client_credentials"},
	}
	body := bytes.NewBufferString(form.Encode())

	resp, err := http.Post("https://api-qa.petfinder.com/v2/oauth2/token/", "application/x-www-form-urlencoded", body)
	if err != nil {
		return "", err
	}

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var token BearerToken
	err = json.Unmarshal(bodyByte, &token)
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}
