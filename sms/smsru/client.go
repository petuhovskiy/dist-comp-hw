package smsru

import "net/http"

const DefaultHost = "https://sms.ru"

type Client struct {
	host  string
	apiID string
	cli   *http.Client
}

func NewClient(host string, apiID string) *Client {
	return &Client{
		host:  host,
		apiID: apiID,
		cli:   http.DefaultClient,
	}
}
