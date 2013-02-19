package main

import (
	"encoding/json"
	"fmt"
	"github.com/jdiez17/go-twitterstream-oauth"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
	"os"
)

func GetOauth() twitterstream.Oauth {
	return twitterstream.NewOauth(Config.TwitterConsumerKey, Config.TwitterConsumerSecret)
}

func store(accessToken *oauth.AccessToken, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := json.Marshal(accessToken)
	if err != nil {
		return err
	}
	file.Write(bytes)

	return nil
}

func retrieve(filename string) (*oauth.AccessToken, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	accessToken := new(oauth.AccessToken)
	json.Unmarshal(bytes, accessToken)

	return accessToken, nil
}

func SaveAccessToken(filename string) {
	Oauth := GetOauth()

	req, err := Oauth.NewAuthenticationRequest()

	if err != nil {
		log.Fatal(err)
	}

	code := ""
	fmt.Println("Authorize at " + req.Url)
	fmt.Println("Gimme your code: ")
	fmt.Scanln(&code)

	accessToken, err := Oauth.GetAccessToken(req.RequestToken, code)
	if err != nil {
		log.Fatal(err)
	}

	err = store(accessToken, filename)
	if err != nil {
		log.Fatal(err)
	}
}
