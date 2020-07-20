package go_onedrive

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	AuthorizationCode = "authorization_code"
	GrantType         = "grant_type"
	ClientId          = "client_id"
	RedirectUri       = "redirect_uri"
	ClientSecret      = "client_secret"
	Code              = "code"

	GraphUrl      = "https://graph.microsoft.com/v1.0/"
	Authorization = "Authrozation"
)

// *****
// GetTokenWithCode returns the response from the Microsoft's server, including the token_type, scope, expiration,
// access_token, refresh_token and id_token.
// *****
func GetTokenWithCode(clientId, redirect, clientSecret, code string) ([]byte, error) {
	data := url.Values{}
	data.Set(GrantType, AuthorizationCode)
	data.Set(ClientId, clientId)
	data.Set(RedirectUri, redirect)
	data.Set(ClientSecret, clientSecret)
	data.Set(Code, code)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://login.microsoftonline.com/common/oauth2/v2.0/token", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Failed to read response's body: " + err.Error())
	}

	return body, nil
}

// *****
// GetAllFiles returns all the files in a folder. If the folderId is an empty string or "root", it returns all the files in the drive's root.
// *****
func GetAllFiles(token, folderId string) ([]byte, error) {
	data := url.Values{}
	data.Set(Authorization, "Bearer "+token)

	client := &http.Client{}

	var fullUrl string
	if len(folderId) != 0 {
		fullUrl = fmt.Sprintf("%s/drive/%s/children", GraphUrl, folderId)
	} else {
		fullUrl = fmt.Sprintf("%s/drive/root/children", GraphUrl)
	}
	req, err := http.NewRequest("POST", fullUrl, strings.NewReader(data.Encode()))
	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Failled to read response's body " + err.Error())
	}

	return body, nil
}