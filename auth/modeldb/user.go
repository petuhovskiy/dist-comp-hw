package modeldb

import (
	"lib/pb"
	"time"
)

type User struct {
	ID           uint
	CreatedAt    time.Time
	Email        string
	Phone        string
	PasswordHash []byte
	Role         pb.AuthRole
}
