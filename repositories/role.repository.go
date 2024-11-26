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
	res, err := roleRepo.dao.ReadAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (roleRepo *RoleRepository) ReadOne(id int64) (*dtos.Role, error) {
	res, err := roleRepo.dao.ReadOne(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (roleRepo *RoleRepository) Create(role dtos.Role) (*dtos.Role, error) {
	res, err := roleRepo.dao.Create(role)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (roleRepo *RoleRepository) Update(role dtos.Role) (*dtos.Role, error) {
	res, err := roleRepo.dao.Update(role)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (roleRepo *RoleRepository) Delete(id int64) (*dtos.Role, error) {
	res, err := roleRepo.dao.Delete(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
