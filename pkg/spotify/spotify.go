package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Spotify struct {
	GrantType    string
	ClientID     string
	ClientSecret string
	TokenUrl     string
	ApiV1        string
	AccessToken  string
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int16  `json:"expires_in"`
}

func (s *Spotify) ShowCredentials() {
	println("GRANT TYPE: " + s.GrantType)
	println("CLIENT ID: " + s.ClientID)
	println("CLIENT SECRET: " + s.ClientSecret)
	println("TOKEN URL: " + s.TokenUrl)
	println("API V1 URL: " + s.ApiV1)
}

func (s *Spotify) Authenticate() {

	data := url.Values{}
	data.Set("grant_type", s.GrantType)
	data.Set("client_id", s.ClientID)
	data.Set("client_secret", s.ClientSecret)

	res, err := http.PostForm(s.TokenUrl, data)

	if err != nil {
		log.Fatal("ERROR AUTHENTICATING:\n", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic("ERROR AUTH RESPONSE BODY")
	}

	if res.StatusCode != 200 {
		panic("Error getting token with status " + fmt.Sprint(res.StatusCode) + " and body:\n" + string(body))
	}

	var result AuthResponse

	json.Unmarshal(body, &result)

	if err != nil {
		panic("ERROR PARSING AUTH RESPONSE:\n" + err.Error())
	}

	s.AccessToken = result.AccessToken
}
