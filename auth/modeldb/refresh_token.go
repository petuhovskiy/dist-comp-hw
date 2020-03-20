package modeldb

import "time"

type RefreshToken struct {
	ID        uint
	CreatedAt time.Time
	UserID    uint
	Token     string
	ExpireAt  time.Time
}
