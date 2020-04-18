package modelapi

import "lib/pb"

type SetRole struct {
	UserID uint        `json:"user_id"`
	Role   pb.AuthRole `json:"role"`
}
