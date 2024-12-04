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
	res, _ := userRepo.dao.ReadAll()
	return res, nil
}

func (userRepo *UserRepository) ReadOne(id int64) (*dtos.User, error) {
	res, _ := userRepo.dao.ReadOne(id)
	return res, nil
}

func (userRepo *UserRepository) Create(user dtos.User) (*dtos.User, error) {
	res, _ := userRepo.dao.Create(user)
	return res, nil
}

func (userRepo *UserRepository) Update(user dtos.User) (*dtos.User, error) {
	res, _ := userRepo.dao.Update(user)
	return res, nil
}

func (userRepo *UserRepository) Delete(id int64) (*dtos.User, error) {
	res, _ := userRepo.dao.Delete(id)
	return res, nil
}
