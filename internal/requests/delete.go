package requests

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func Remove(wg *sync.WaitGroup, ch chan string, id int) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	c := http.Client{}

	urlUser := "http://127.0.0.1:8080/user"
	authEncodedCreds := "U29tZVVlcjpzaW1wbGVzdFBhc3N3b3Jk"

	req, err := http.NewRequest("DELETE", urlUser, nil)
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

	ch <- fmt.Sprintf("Response from the GET request by all users: %v\n", string(body))
}
