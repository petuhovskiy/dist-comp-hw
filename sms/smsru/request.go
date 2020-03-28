package smsru

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/url"
)

func (c *Client) doPost(path string, values url.Values, res interface{}) error {
	values.Add("api_id", c.apiID)
	values.Add("json", "1")
	url := c.host + path + "?" + values.Encode()

	resp, err := c.cli.Post(url, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	log.WithField("body", string(b)).Info("response from smsru")
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, res)
	return err
}