package services

import (
	"github.com/ssjlee93/fitworks-data-user/models"
	"github.com/ssjlee93/fitworks-data-user/repositories"
)

type RoleService struct {
	r repositories.RoleRepository
}

func NewRoleService(r repositories.RoleRepository) *RoleService {
	return &RoleService{r: r}
}

func (roleService *RoleService) ReadAll() ([]models.Role, error) {
	res, _ := roleService.r.ReadAll()
	return res, nil
}

func (roleService *RoleService) ReadOne(id int64) (*models.Role, error) {
	res, _ := roleService.r.ReadOne(id)
	return res, nil
}

func (roleService *RoleService) Create(role models.Role) (*models.Role, error) {
	res, _ := roleService.r.Create(role)
	return res, nil
}

func (roleService *RoleService) Update(role models.Role) (*models.Role, error) {
	res, _ := roleService.r.Update(role)
	return res, nil
}

func (roleService *RoleService) Delete(id int64) (*models.Role, error) {
	res, _ := roleService.r.Delete(id)
	return res, nil
}
