package repositories

import (
	"github.com/ssjlee93/fitworks-data-user/daos"
	"github.com/ssjlee93/fitworks-data-user/dtos"
)

type UserRepository struct {
	dao daos.UserDAOImpl
}

func NewUserRepository(dao daos.UserDAOImpl) *UserRepository {
	return &UserRepository{dao: dao}
}

func (userRepo *UserRepository) ReadAll() ([]dtos.User, error) {
	res, err := userRepo.dao.ReadAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (userRepo *UserRepository) ReadOne(id int64) (*dtos.User, error) {
	res, err := userRepo.dao.ReadOne(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (userRepo *UserRepository) Create(user dtos.User) (*dtos.User, error) {
	res, err := userRepo.dao.Create(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (userRepo *UserRepository) Update(user dtos.User) (*dtos.User, error) {
	res, err := userRepo.dao.Update(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (userRepo *UserRepository) Delete(id int64) (*dtos.User, error) {
	res, err := userRepo.dao.Delete(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
