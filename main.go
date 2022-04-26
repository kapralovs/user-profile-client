package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	urlUser := "http://127.0.0.1:8080/user"
	urlUserID := "http://127.0.0.1:8080/user/"
	authEncodedCreds := "U29tZVVlcjpzaW1wbGVzdFBhc3N3b3Jk"
	client := http.Client{}
	for i := 1; i <= 3; i++ {
		go func() {
			req, err := http.NewRequest("GET", urlUser, nil)
			if err != nil {
				log.Println(err)
			}

			req.Header.Add("Authorization", "Basic "+authEncodedCreds)
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}

			fmt.Printf("Response from the get request by all users: %v\n", string(body))
		}()
		go func(intId int) {
			id := strconv.Itoa(intId)
			buffer := bytes.NewBuffer([]byte("{\"id\": \"534\",\"email\": \"anothername@email.com\",\"username\": \"anotherUserName\",\"password\": \"ultraPassword\",\"is_admin\": true}"))
			parts := []string{
				urlUserID,
				id,
			}
			url := strings.Join(parts, "")
			req, err := http.NewRequest("POST", url, buffer)
			if err != nil {
				log.Println(err)
			}

			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", "Basic"+authEncodedCreds)
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}

			fmt.Printf("Response from the POST request by user by id: %v\n", string(body))
		}(i)
	}
}
