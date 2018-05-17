package main

import (
	"fmt"
	"net/http"
)

func UrlValid(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return true
}

func main() {
	fmt.Println(UrlValid("http://www.basssidu.com"))
}
