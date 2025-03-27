package services

import (
	"github.com/ssjlee93/fitworks-data-user/models"
	"github.com/ssjlee93/fitworks-data-user/repositories"
)

type RoleRepository struct {
	dao repositories.RoleDAOImpl
}

func NewRoleRepository(dao repositories.RoleDAOImpl) *RoleRepository {
	return &RoleRepository{dao: dao}
}

func (roleRepo *RoleRepository) ReadAll() ([]models.Role, error) {
	res, _ := roleRepo.dao.ReadAll()
	return res, nil
}

func (roleRepo *RoleRepository) ReadOne(id int64) (*models.Role, error) {
	res, _ := roleRepo.dao.ReadOne(id)
	return res, nil
}

func (roleRepo *RoleRepository) Create(role models.Role) (*models.Role, error) {
	res, _ := roleRepo.dao.Create(role)
	return res, nil
}

func (roleRepo *RoleRepository) Update(role models.Role) (*models.Role, error) {
	res, _ := roleRepo.dao.Update(role)
	return res, nil
}

func (roleRepo *RoleRepository) Delete(id int64) (*models.Role, error) {
	res, _ := roleRepo.dao.Delete(id)
	return res, nil
}
