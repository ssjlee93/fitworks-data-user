package services

import (
	"github.com/ssjlee93/fitworks-data-user/models"
	"github.com/ssjlee93/fitworks-data-user/repositories"
	"log"
)

type UserService struct {
	r repositories.Repository[models.User]
}

func NewUserService(r repositories.UserRepository) *UserService {
	return &UserService{r: &r}
}

func (userService *UserService) ReadAll() ([]models.User, error) {
	log.Println("| - UserService.ReadAll")
	res, _ := userService.r.ReadAll()
	return res, nil
}

func (userService *UserService) ReadOne(id int64) (*models.User, error) {
	log.Println("| - UserService.ReadOne")
	res, _ := userService.r.ReadOne(id)
	return res, nil
}

func (userService *UserService) Create(user models.User) error {
	log.Println("| - UserService.Create")
	err := userService.r.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (userService *UserService) Update(user models.User) error {
	log.Println("| - UserService.Update")
	err := userService.r.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (userService *UserService) Delete(id int64) error {
	log.Println("| - UserService.Delete")
	err := userService.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
