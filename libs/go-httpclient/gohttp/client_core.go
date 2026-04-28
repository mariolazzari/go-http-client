package gohttp

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(method, url string, headers http.Header, body any) (*http.Response, error) {
	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.New("unable to create a request")
	}

	request.Header = c.getRequestHeaders(headers)

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
