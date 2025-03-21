package repositories

import (
	"github.com/ssjlee93/fitworks-data-user/daos"
	"github.com/ssjlee93/fitworks-data-user/dtos"
	"log"
)

type UserRepository struct {
	dao daos.Dao[dtos.User]
}

func NewUserRepository(dao daos.UserDAOImpl) *UserRepository {
	return &UserRepository{dao: &dao}
}

func (userRepo *UserRepository) ReadAll() ([]dtos.User, error) {
	log.Println("| - - UserRepository.ReadAll")
	res, _ := userRepo.dao.ReadAll()
	return res, nil
}

func (userRepo *UserRepository) ReadOne(id int64) (*dtos.User, error) {
	log.Println("| - - UserRepository.ReadOne")
	res, _ := userRepo.dao.ReadOne(id)
	return res, nil
}

func (userRepo *UserRepository) Create(user dtos.User) error {
	log.Println("| - - UserRepository.Create")
	err := userRepo.dao.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (userRepo *UserRepository) Update(user dtos.User) error {
	log.Println("| - - UserRepository.Update")
	err := userRepo.dao.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (userRepo *UserRepository) Delete(id int64) error {
	log.Println("| - - UserRepository.Delete")
	err := userRepo.dao.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
