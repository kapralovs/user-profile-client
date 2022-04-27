package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Profile struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func main() {
	urlUserID := "http://127.0.0.1:8080/user/"
	authEncodedCreds := "U29tZVVlcjpzaW1wbGVzdFBhc3N3b3Jk"
	client := http.Client{}

	profile := Profile{
		ID:       "2",
		Email:    "anotherEmail@email.ru",
		Username: "awesomeNickName",
		Password: "NewUserPassword",
		IsAdmin:  false,
	}

	asBytes, err := json.Marshal(profile)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("json as string: ", string(asBytes))
	// id := "4"
	// buffer := bytes.NewBuffer(asBytes)
	parts := []string{
		urlUserID,
		profile.ID,
	}
	url := strings.Join(parts, "")
	req, err := http.NewRequest("POST", url, strings.NewReader(string(asBytes)))
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+authEncodedCreds)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Response from the POST request by user by id: %v\n", string(body))
}
