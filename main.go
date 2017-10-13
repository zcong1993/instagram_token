package main

import (
	"os"
	"log"
)

const (
	TokenUrl    = "https://api.instagram.com/oauth/access_token"
	RedirectUrl = "http://localhost:7080/callback"
	AppName     = "instagram access token"
)

func init() {
	ClientID := os.Getenv("CLIENT_ID")
	ClientSecret := os.Getenv("CLIENT_SECRET")
	if ClientID == "" || ClientSecret == "" {
		log.Fatal("env `ClientID` and `ClientSecret` are required!")
	}
}

func main() {
	ClientID := os.Getenv("CLIENT_ID")
	ClientSecret := os.Getenv("CLIENT_SECRET")
	Run(ClientID, ClientSecret)
}
