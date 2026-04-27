package main

import (
	"fmt"
	"github/mariolazzari/go-http-client/libs/go-httpclient/gohttp"
	"io"
	"log"
)

func main() {
	client := gohttp.New()
	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
