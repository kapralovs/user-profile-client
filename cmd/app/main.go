package main

import (
	"client/internal/requests"
	"fmt"
	"sync"
)

type Profile struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

var wg sync.WaitGroup

func main() {
	counter := 10

	goroutinesCount := counter * 3

	ch := make(chan string, goroutinesCount)

	wg.Add(goroutinesCount)
	wgPntr := &wg
	for i := 1; i <= counter; i++ {
		go requests.Get(i, wgPntr, ch)
		go requests.Create(wgPntr, ch, i)
		go requests.GetByID(wgPntr, ch, 2)
		// go requests.Remove(wgPntr, ch, i-1)
		go requests.Edit(wgPntr, ch, i)
	}
	wg.Wait()

	for j := 1; j <= goroutinesCount; j++ {
		fmt.Println(<-ch)
	}
}
