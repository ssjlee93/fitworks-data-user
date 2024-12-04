package repositories

import (
	"github.com/ssjlee93/fitworks-data-user/daos"
	"github.com/ssjlee93/fitworks-data-user/dtos"
)

type RoleRepository struct {
	dao daos.RoleDAOImpl
}

func NewRoleRepository(dao daos.RoleDAOImpl) *RoleRepository {
	return &RoleRepository{dao: dao}
}

func (roleRepo *RoleRepository) ReadAll() ([]dtos.Role, error) {
	res, _ := roleRepo.dao.ReadAll()
	return res, nil
}

func (roleRepo *RoleRepository) ReadOne(id int64) (*dtos.Role, error) {
	res, _ := roleRepo.dao.ReadOne(id)
	return res, nil
}

func (roleRepo *RoleRepository) Create(role dtos.Role) (*dtos.Role, error) {
	res, _ := roleRepo.dao.Create(role)
	return res, nil
}

func (roleRepo *RoleRepository) Update(role dtos.Role) (*dtos.Role, error) {
	res, _ := roleRepo.dao.Update(role)
	return res, nil
}

func (roleRepo *RoleRepository) Delete(id int64) (*dtos.Role, error) {
	res, _ := roleRepo.dao.Delete(id)
	return res, nil
}
