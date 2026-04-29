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
