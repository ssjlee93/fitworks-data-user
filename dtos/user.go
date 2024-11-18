package dtos

import (
	"fmt"
	"time"
)

type User struct {
	UserID    int64   `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Google    *string `json:"google"`
	Apple     *string `json:"apple"`
	RoleID    int64   `json:"role_id"`
	Role      Role    `json:"role"`
	TrainerID *int64  `json:"trainer_id"`
	Trainer   *User   `json:"trainer"`
	Created   time.Time
	Updated   time.Time
}

func (user *User) PrintUser() {
	fmt.Println("=== Print User ===")
	fmt.Println("user_id: ", user.UserID)
	fmt.Println("first_name: ", user.FirstName)
	fmt.Println("last_name: ", user.LastName)
	fmt.Println("google: ", user.Google)
	fmt.Println("apple: ", user.Apple)
	fmt.Println("role_id: ", user.RoleID)
	fmt.Println("role: ", user.Role)
	fmt.Println("trainer_id: ", user.TrainerID)
	fmt.Println("created: ", user.Created)
	fmt.Println("updated: ", user.Updated)
}
