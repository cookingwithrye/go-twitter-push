package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type MatchPattern struct {
	From, MatchText string
}

type ConfigValues struct {
	PushoverAPIKey, PushoverUser              string
	TwitterConsumerKey, TwitterConsumerSecret string
	NotificationTitle, NotificationBody       string
	Patterns                                  []MatchPattern
}

func ReadConfig(filename string) (*ConfigValues, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	config := new(ConfigValues)
	err = json.Unmarshal(bytes, config)
	return config, err
}
