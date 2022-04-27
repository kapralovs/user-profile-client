package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	c := http.Client{}

	urlUser := "http://127.0.0.1:8080/user"
	authEncodedCreds := "U29tZVVlcjpzaW1wbGVzdFBhc3N3b3Jk"

	req, err := http.NewRequest("GET", urlUser, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Authorization", "Basic "+authEncodedCreds)
	resp, err := c.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Response from the get request by all users: %v\n", string(body))
}
