# How to develop a productive HTTP client in Golang (Go)

[Github](https://github.com/federicoleon/go-httpclient)

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

```go
package gohttp

import "net/http"

type HttpClient interface {
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body any) (*http.Response, error)
	Patch(url string, headers http.Header, body any) (*http.Response, error)
	Put(url string, headers http.Header, body any) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

func New() HttpClient {
	return &httpClient{}
}

type httpClient struct {
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)

}

func (c *httpClient) Post(url string, headers http.Header, body any) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body any) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body any) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
```

### Custom headers

```go
package gohttp

import "net/http"

type HttpClient interface {
	SetHeaders(headers http.Header)
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body any) (*http.Response, error)
	Patch(url string, headers http.Header, body any) (*http.Response, error)
	Put(url string, headers http.Header, body any) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

type httpClient struct {
	Header http.Header
}

func New() HttpClient {
	return &httpClient{}
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.Header = headers
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)

}

func (c *httpClient) Post(url string, headers http.Header, body any) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body any) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body any) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
```

### Request body

```go
package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
)

func (c *httpClient) do(method, url string, headers http.Header, body any) (*http.Response, error) {
	client := http.Client{}

	fullHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create a request")
	}

	request.Header = fullHeaders

	return client.Do(request)
}

func (c *httpClient) getRequestHeaders(requestHeader http.Header) http.Header {
	result := make(http.Header)

	// Common headers
	for header, value := range c.Header {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Custom headers
	for header, value := range requestHeader {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	return result
}

func (c *httpClient) getRequestBody(contentType string, body any) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}
```

### Testing

- Initialization
- Execution
- Validation

```go
package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("UIser-Agent", "http-client")
	client.Headers = commonHeaders

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "abc1234")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(finalHeaders) != 3 {
		t.Error("3 errors expected")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type")
	}

	if finalHeaders.Get("X-Request-Id") != "abc1234" {
		t.Error("invalid rewuest id")
	}
}

func TestGetRequestBody(t *testing.T) {

	// Init
	client := httpClient{}

	t.Run("BodyNil", func(t *testing.T) {
		// Exec
		body, err := client.getRequestBody("", nil)

		// Valid
		if err != nil {
			t.Error("no error expected with nil body")
		}
		if body != nil {
			t.Error("no body expected with nil body")
		}
	})

	t.Run("BodyJson", func(t *testing.T) {
		// Exec
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)

		// Valid
		if err != nil {
			t.Error("no error expected when marshaling")
		}
		if string(body) != `["one","two"]` {
			t.Error("invalid json body")
		}
	})

	t.Run("BodyXml", func(t *testing.T) {

	})

	t.Run("Default", func(t *testing.T) {

	})
}
```

## Publishing
