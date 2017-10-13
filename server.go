package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"bytes"
	"net/url"
	"io/ioutil"
	"github.com/buger/jsonparser"
)

//type User struct {
//	ID string `decoder:"id"`
//	Username string `decoder:"username"`
//	ProfilePicture string `decoder:"profile_picture"`
//	FullName string `decoder:"full_name"`
//	Bio string `decoder:"bio"`
//	Website string `decoder:"website"`
//	IsBusiness bool `decoder:"is_business"`
//}
//
//type Resp struct {
//	AccessToken string `decoder:"access_token"`
//	User User `decoder:"user"`
//}

func Post(apiUrl string, form map[string]string, headers map[string]string)([]byte, error) {
	client := &http.Client{}
	formData := url.Values{}
	for k, v := range form {
		formData.Set(k, v)
	}
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	token, _, _, err := jsonparser.Get(d, "access_token")
	if err != nil {
		return nil, err
	}
	return token, nil
}

func Run(ClientID, ClientSecret string) {
	link := fmt.Sprintf("https://api.instagram.com/oauth/authorize/?client_id=%s&redirect_uri=%s&response_type=code", ClientID, RedirectUrl)
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"link": link,
			"name": AppName,
		})
	})
	router.GET("/callback", func(c *gin.Context) {
		code := c.Query("code")
		//params := &Params{ClientID, ClientSecret, "authorization_code", RedirectUrl, code}
		formData := map[string]string{
			"client_id": ClientID,
			"client_secret": ClientSecret,
			"grant_type": "authorization_code",
			"redirect_uri": RedirectUrl,
			"code": code,
		}
		var msg string
		token, err := Post(TokenUrl, formData, map[string]string{"Content-Type": "application/x-www-form-urlencoded"})
		if err != nil {
			msg = "Oops! An error occurred!"
		}
		msg = string(token)
		fmt.Println(msg)
		c.HTML(http.StatusOK, "show.tmpl", gin.H{
			"token": msg,
			"name": AppName,
		})
	})
	router.Run(":7080")
}