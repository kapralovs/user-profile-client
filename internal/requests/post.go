package requests

import (
	"client/internal/data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Edit(wg *sync.WaitGroup, ch chan string, id int) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	urlUserID := "http://127.0.0.1:8080/user/"
	authEncodedCreds := "U29tZVVlcjpzaW1wbGVzdFBhc3N3b3Jk"
	client := http.Client{}

	strID := strconv.Itoa(id)
	newID := "20" + strID
	uName := "justEditedBot" + strID
	email := uName + "@email.com"
	psswrd := uName + "EditedPsswd"

	profile := &data.Profile{
		ID:       newID,
		Email:    email,
		Username: uName,
		Password: psswrd,
		IsAdmin:  true,
	}

	asBytes, err := json.Marshal(profile)
	if err != nil {
		log.Println(err)
	}

	parts := []string{
		urlUserID,
		strID,
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

	ch <- fmt.Sprintf("Goroutine %v.Response from the POST (edit) request by user by ID (%v): %v\n", id, strID, string(body))
}

func Create(wg *sync.WaitGroup, ch chan string, id int) {
	defer wg.Done()
	url := "http://127.0.0.1:8080/user"
	authEncodedCreds := "U29tZVVlcjpzaW1wbGVzdFBhc3N3b3Jk"
	client := http.Client{}

	strID := strconv.Itoa(id + 100)
	uName := "justBot" + strID
	email := uName + "@email.com"
	psswrd := uName + "Psswd"

	profile := &data.Profile{
		ID:       strID,
		Email:    email,
		Username: uName,
		Password: psswrd,
		IsAdmin:  true,
	}

	asBytes, err := json.Marshal(profile)
	if err != nil {
		log.Println(err)
	}

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

	ch <- fmt.Sprintf("Goroutine %v.Response from the POST (create) request by user by ID (%v): %v\n", id, strID, string(body))
}
