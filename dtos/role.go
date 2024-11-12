package dtos

import "time"

type Role struct {
	RoleID  uint
	Role    string
	Created time.Time
	Updated time.Time
}
