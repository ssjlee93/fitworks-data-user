package daos

import "github.com/ssjlee93/fitworks-data-user/dtos"

type RoleDao interface {
	Create()
	ReadOne() (dtos.Role, error)
	ReadAll() ([]dtos.Role, error)
	Update()
	Delete()
}
