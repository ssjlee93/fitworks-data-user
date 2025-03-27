package services

import (
	"github.com/ssjlee93/fitworks-data-user/models"
	"github.com/ssjlee93/fitworks-data-user/repositories"
	"log"
)

type UserRepository struct {
	dao repositories.Dao[models.User]
}

func NewUserRepository(dao repositories.UserDAOImpl) *UserRepository {
	return &UserRepository{dao: &dao}
}

func (userRepo *UserRepository) ReadAll() ([]models.User, error) {
	log.Println("| - - UserRepository.ReadAll")
	res, _ := userRepo.dao.ReadAll()
	return res, nil
}

func (userRepo *UserRepository) ReadOne(id int64) (*models.User, error) {
	log.Println("| - - UserRepository.ReadOne")
	res, _ := userRepo.dao.ReadOne(id)
	return res, nil
}

func (userRepo *UserRepository) Create(user models.User) error {
	log.Println("| - - UserRepository.Create")
	err := userRepo.dao.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (userRepo *UserRepository) Update(user models.User) error {
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
