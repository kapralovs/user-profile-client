package requests

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func Get(counter int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
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

	ch <- fmt.Sprintf("Goroutine %v.Response from the GET request by all users: %v\n", counter, string(body))
}

func GetByID(wg *sync.WaitGroup, ch chan string, id int) {
	defer wg.Done()
	c := http.Client{}

	urlUser := "http://127.0.0.1:8080/user/"
	authEncodedCreds := "U29tZVVlcjpzaW1wbGVzdFBhc3N3b3Jk"

	strID := strconv.Itoa(id)

	parts := []string{
		urlUser,
		strID,
	}

	url := strings.Join(parts, "")
	req, err := http.NewRequest("GET", url, nil)
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

	ch <- fmt.Sprintf("Goroutine %v.Response from the GET (by ID) request by ID (%v): %v\n", id, strID, string(body))
}
