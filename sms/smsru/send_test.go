package auth

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_Send(t *testing.T) {
	apiID := os.Getenv("TEST_API_ID")
	if apiID == "" {
		t.Skip("TEST_API_ID is not set")
	}

	cli := NewClient(DefaultHost, apiID)
	resp, err := cli.Send(os.Getenv("TEST_SMS_RECIPIENT"), "sample sms message")
	assert.Nil(t, err)
	spew.Dump(resp)
}