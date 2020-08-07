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
	GrantType         = "grant_type"
	AuthorizationCode = "authorization_code"
	ClientId          = "client_id"
	RedirectUri       = "redirect_uri"
	ClientSecret      = "client_secret"
	RefreshToken      = "refresh_token"
	Code              = "code"

	GraphUrl      = "https://graph.microsoft.com/v1.0"
	Authorization = "Authorization"
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
	if err != nil {
		return nil, errors.New("Failled assemble request" + err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Failled exec request" + err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Failed to read response's body: " + err.Error())
	}

	return body, nil
}

// *****
// GetTokenWithRefreshToken returns the response from the Microsoft's server, including the token_type, scope, expiration,
// access_token, refresh_token and id_token.
// *****
func GetTokenWithRefreshToken(refreshToken, clientId, redirect, clientSecret string) ([]byte, error) {
	data := url.Values{}
	data.Set(GrantType, RefreshToken)
	data.Set(ClientId, clientId)
	data.Set(RedirectUri, redirect)
	data.Set(ClientSecret, clientSecret)
	data.Set(RefreshToken, refreshToken)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://login.microsoftonline.com/common/oauth2/v2.0/token", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, errors.New("Failled assemble request" + err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Failled exec request" + err.Error())
	}
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
	client := &http.Client{}

	var fullUrl string
	if len(folderId) != 0 {
		fullUrl = fmt.Sprintf("%s/me/drive/%s/children", GraphUrl, folderId)
	} else {
		fullUrl = fmt.Sprintf("%s/me/drive/root/children", GraphUrl)
	}
	req, err := http.NewRequest("GET", fullUrl, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	if err != nil {
		return nil, errors.New("Failled assemble request" + err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Failled to exec request " + err.Error())
	}

	if res.StatusCode == 401 {
		return nil, errors.New("Expired token")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Failled to read response's body " + err.Error())
	}

	return body, nil
}

// *****
// UploadFile's limit is 4mb. Returns the Onedrive object of the file.
// *****
func UploadFile(token, file, folderID string) ([]byte, error) {
	client := &http.Client{}

	var fullUrl string
	if len(folderID) != 0 {
		fullUrl = fmt.Sprintf("%s/me/drive/items/%s/children", GraphUrl, folderID)
	} else {
		fullUrl = fmt.Sprintf("%s/me/drive/root/children", GraphUrl)
	}
	req, err := http.NewRequest("PUT", fullUrl, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	if err != nil {
		return nil, errors.New("Failled assemble request" + err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Failled to exec request " + err.Error())
	}

	if res.StatusCode == 401 {
		return nil, errors.New("Expired token")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Failled to read response's body " + err.Error())
	}

	return body, nil
}

