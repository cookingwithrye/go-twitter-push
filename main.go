package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/jdiez17/go-twitterstream-oauth"
	"strings"
	"text/template"
)

var Config *ConfigValues = nil
var NotificationTemplate *template.Template = nil

func main() {
	getAccessToken := flag.Bool("get-access-token", false, "If true, save the access token.")
	filename := flag.String("filename", "tokens.json", "Location to store the access tokens")
	config := flag.String("config", "config.json", "Config file")

	var err error

	Config, err = ReadConfig(*config)
	if err != nil {
		fmt.Println("Your config file is invalid. Error: ", err)
		return
	}
	NotificationTemplate, err = template.New("notification").Parse(Config.NotificationBody)
	if err != nil {
		fmt.Println("Your NotificationBody is malformed: ", err)
		return
	}

	flag.Parse()
	if *getAccessToken {
		SaveAccessToken(*filename)
		return
	}

	accessToken, err := retrieve(*filename)
	if err != nil {
		fmt.Println("You must store the access token with --get-access-token")
		return
	}
	tweets, err := twitterstream.DoUserStream(GetOauth(), accessToken)
	if err != nil {
		fmt.Println("Error openng the twitter stream.")
		panic(err)
	}

	var buf bytes.Buffer
	for tw := range tweets {
		fmt.Println(tw)
		for _, p := range Config.Patterns {
			if p.From == "*" || p.From == tw.User.Screen_name {
				if p.MatchText == "*" || strings.Contains(tw.Text, p.MatchText) {
					NotificationTemplate.Execute(&buf, tw)
					Notify(buf.String())
				}
			}
		}
	}
}
