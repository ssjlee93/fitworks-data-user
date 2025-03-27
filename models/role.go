package models

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
	fmt.Println("=== Printing role ===")
	fmt.Println("role_id: ", role.RoleID)
	fmt.Println("role: ", role.Role)
	fmt.Println("created: ", role.Created)
	fmt.Println("updated: ", role.Updated)
}
