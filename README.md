# How to develop a productive HTTP client in Golang (Go)

## HTTP calls

### Basic GET client

```go
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
```

### HTTP Server

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("GET /", rootHandler)
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	fmt.Fprintf(w, "Ciao Mario!")
}
```

## Working on the core

### Go modules

- create github repo
- .gitignore
- go mod init
- go mod tidy
- no main

### Go basics

```go
package gohttp

type HttpClient interface {
	Get()
	Post()
	Patch()
	Put()
	Delete()
}

func New() HttpClient {
	return &httpClient{}
}

type httpClient struct {
}

func (c *httpClient) Get() {

}

func (c *httpClient) Post() {

}

func (c *httpClient) Patch() {

}

func (c *httpClient) Put() {

}

func (c *httpClient) Delete() {

}
```

### Basic behavior
