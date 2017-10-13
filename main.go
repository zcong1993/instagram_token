package main

import (
	"os"
)

const (
	TokenUrl = "https://api.instagram.com/oauth/access_token"
	RedirectUrl = "http://localhost:7080/callback"
	AppName = "instagram access token"
)

func main() {
	ClientID := os.Getenv("CLIENT_ID")
	ClientSecret := os.Getenv("CLIENT_SECRET")
	Run(ClientID, ClientSecret)
}