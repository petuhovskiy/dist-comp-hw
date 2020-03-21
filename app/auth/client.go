package auth

import "net/http"

type Client struct {
	url string
	cli *http.Client
}

func NewClient(url string) *Client {
	return &Client{
		url: url,
		cli: http.DefaultClient,
	}
}