package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://vk.com/")
	if err != nil {
		log.Fatal(nil)
	}

}
