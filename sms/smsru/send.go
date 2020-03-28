package smsru

import (
	"errors"
	"net/url"
)

type SendResponse struct {
	Status     string             `json:"status"`
	StatusCode int                `json:"status_code"`
	SmsInfo    map[string]NumInfo `json:"sms"`
	Balance    float64            `json:"balance"`
}

type NumInfo struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	SmsID      string `json:"sms_id"`
	StatusText string `json:"status_text"`
}

func (c *Client) Send(to string, msg string) (SendResponse, error) {
	req := make(url.Values)
	req.Add("to", to)
	req.Add("msg", msg)

	var resp SendResponse
	err := c.doPost("/sms/send", req, &resp)
	if err == nil && resp.Status != "OK" {
		err = errors.New("resp status is not ok")
	}

	return resp, err
}