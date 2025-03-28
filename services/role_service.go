package services

import (
	"github.com/ssjlee93/fitworks-data-user/models"
	"github.com/ssjlee93/fitworks-data-user/repositories"
)

type RoleService struct {
	dao repositories.RoleDAOImpl
}

func NewRoleService(dao repositories.RoleDAOImpl) *RoleService {
	return &RoleService{dao: dao}
}

func (roleRepo *RoleService) ReadAll() ([]models.Role, error) {
	res, _ := roleRepo.dao.ReadAll()
	return res, nil
}

func (roleRepo *RoleService) ReadOne(id int64) (*models.Role, error) {
	res, _ := roleRepo.dao.ReadOne(id)
	return res, nil
}

func (roleRepo *RoleService) Create(role models.Role) (*models.Role, error) {
	res, _ := roleRepo.dao.Create(role)
	return res, nil
}

func (roleRepo *RoleService) Update(role models.Role) (*models.Role, error) {
	res, _ := roleRepo.dao.Update(role)
	return res, nil
}

func (roleRepo *RoleService) Delete(id int64) (*models.Role, error) {
	res, _ := roleRepo.dao.Delete(id)
	return res, nil
}
