package go_onedrive

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetTokenWithCode(clientId, redirect, clientSecret, code string) ([]byte, error) {
	requestBody, err := json.Marshal(map[string]string{
		"client_id": clientId,
		"redirect_uri": redirect,
		"client_secret": clientSecret,
		"code": code,
		"grant_type":"authorization_code",
	})
	if err != nil {
		return nil, errors.New("Couldn't marshal JSON: " + err.Error())
	}
	req, err := http.Post("https://login.microsoftonline.com/common/oauth2/v2.0/token", "x-www-form-urlencoded", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("Failled to post: " + err.Error())
	}

	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, errors.New("Failed to read req's body")
	}
	return body, nil
}



