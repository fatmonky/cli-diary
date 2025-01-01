package main

import (
	"fmt"
	"net/http"
)

func main() {
	c := make(chan string)
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
	}
	for _, link := range links {
		go checkLink(link, c)

	}
	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
		default:
			break
		}
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link + " might not be working!"
		return
	}
	c <- link + " is up!"
}
