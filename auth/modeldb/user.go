package modeldb

import "time"

type User struct {
	ID           uint
	CreatedAt    time.Time
	Email        string
	Phone        string
	PasswordHash []byte
}
