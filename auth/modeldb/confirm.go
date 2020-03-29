package modeldb

import "time"

type ConfirmType string

var (
	ConfirmSms   ConfirmType = "sms"
	ConfirmEmail ConfirmType = "email"
)

type Confirm struct {
	Link     string      // link used for confirmation
	UserID   uint        // user id
	Type     ConfirmType // sms or email
	Subject  string      // corresponding phone number or email address
	ExpireAt time.Time   // expiration time
}
