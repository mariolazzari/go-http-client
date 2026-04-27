package gohttpclient

import "github/mariolazzari/go-http-client/libs/go-httpclient/gohttp"

func exampleUsage() {
	client := gohttp.New()
	client.Get()
}
