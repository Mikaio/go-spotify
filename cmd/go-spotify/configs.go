package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type spotifyConfigs struct {
	ClientID     string
	ClientSecret string
	GrantType    string
	TokenUrl     string
	ApiV1        string
}

var SpotifyConfigs = spotifyConfigs{}

func init() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env")
	}

	SpotifyConfigs.ClientID = os.Getenv("SPOTIFY_CLIENT_ID")
	SpotifyConfigs.ClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
	SpotifyConfigs.GrantType = os.Getenv("SPOTIFY_GRANT_TYPE")
	SpotifyConfigs.TokenUrl = os.Getenv("SPOTIFY_TOKEN_URL")
	SpotifyConfigs.ApiV1 = os.Getenv("SPOTIFY_API_V1")
}
