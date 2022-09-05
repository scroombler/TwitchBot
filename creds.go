package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Cred struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	AppToken     string `json:"appToken"`
}

func LoadCreds(path string) (*Cred, error) {
	var result Cred

	bs, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading credentials:", err)
		return nil, err
	}

	err = json.Unmarshal(bs, &result)
	if err != nil {
		fmt.Println("Error parsing credentials:", err)
		return nil, err
	}

	return &result, nil
}
