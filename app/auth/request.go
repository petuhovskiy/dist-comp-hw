package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrResponse struct {
	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e ErrResponse) Error() string {
	return fmt.Sprintf("status=%v, code=%v, error=%v", e.StatusText, e.AppCode, e.ErrorText)
}

func (c *Client) doPost(path string, req interface{}, res interface{}) error {
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := c.cli.Post(c.url+path, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var customErr ErrResponse
		err = json.NewDecoder(resp.Body).Decode(&customErr)
		if err != nil {
			return err
		}
		return customErr
	}

	err = json.NewDecoder(resp.Body).Decode(res)
	return err
}