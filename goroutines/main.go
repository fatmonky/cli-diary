package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
	}
	wg.Add(3)
	for _, link := range links {
		go checkLink(link)
	}
	wg.Wait()
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might not be working!")
	}
	fmt.Println(link, "is up!")
	defer wg.Done()
}
