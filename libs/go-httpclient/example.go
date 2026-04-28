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
	createUser(User{
		FirstName: "Mario",
		LastName:  "Lazzari",
	})
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getUrls() {
	headers := make(http.Header)
	headers.Set("Authorization", "")

	response, err := githubClient.Get("https://api.github.com", headers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}

func createUser(user User) {

	headers := make(http.Header)
	headers.Set("Authorization", "")

	response, err := githubClient.Post("https://api.github.com", headers, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
