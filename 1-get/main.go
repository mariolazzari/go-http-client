package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "https://api.github.com"

	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("Accept", "application/json")

	// res, err := client.Get(url)
	res, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.StatusCode)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Println(string(body))
}
