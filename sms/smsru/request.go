package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
)

func (c *Client) doPost(path string, values url.Values, res interface{}) error {
	values.Add("api_id", c.apiID)
	url := c.host + path + "?" + values.Encode()

	resp, err := c.cli.Post(url, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	log.Printf("Response from smsru: %s", b)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, res)
	return err
}