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
