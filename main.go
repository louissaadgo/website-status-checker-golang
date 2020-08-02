package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
		"http://golang.org",
	}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		go func(l string) {
			time.Sleep(4 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Oops, Something Went Wrong!")
		c <- link
		return
	}
	fmt.Println(link, "Is Up!")
	c <- link
}
