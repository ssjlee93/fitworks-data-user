package dtos

import "fmt"

type User struct {
	UserID    int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Google    string `json:"google"`
	Apple     string `json:"apple"`
	Role      Role   `json:"role"`
	Trainer   *User  `json:"trainer"`
}

func (user *User) PrintUser() {
	fmt.Println(user.UserID)
	fmt.Println(user.FirstName)
	fmt.Println(user.LastName)
	fmt.Println(user.Google)
	fmt.Println(user.Apple)
	fmt.Println(user.Role.Role)
	fmt.Println(user.Trainer)
}
