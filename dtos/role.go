package dtos

import (
	"fmt"
	"time"
)

type Role struct {
	RoleID  int64
	Role    string
	Created time.Time
	Updated time.Time
}

// PrintRole prints every attribute in Role
// TODO delete this in production
func (role *Role) PrintRole() {
	fmt.Println(role.RoleID)
	fmt.Println(role.Role)
	fmt.Println(role.Created)
	fmt.Println(role.Updated)
}
