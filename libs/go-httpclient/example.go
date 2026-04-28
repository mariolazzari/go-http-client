package main

import (
	"fmt"
	"github/mariolazzari/go-http-client/libs/go-httpclient/gohttp"
	"io"
	"log"
	"net/http"
)

var githubClient = gohttp.New()

func main() {
	getUrls()
}

func getUrls() {

	headers := make(http.Header)
	headers.Set("Authorization", "Bearer 123abc")

	response, err := githubClient.Get("https://api.github.com", headers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
