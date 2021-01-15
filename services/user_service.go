package services

import (
	"fmt"

	"github.com/twisty/clean-arch/imp"
	"github.com/twisty/clean-arch/models"
)

type userService struct {
	userRepo imp.UserRepository
}

func NewService(repo imp.UserRepository) imp.UserService {
	return &userService{userRepo: repo}
}

func (userService *userService) GetAllUser() (users []models.User, err error) {
	var user []models.User
	handle := userService.userRepo.GetAllUser(&user)
	if handle != nil {
		fmt.Println("Error")
	}
	return user, handle
}

func (userService *userService) GetUserById(id string) (users models.User, err error) {
	var user models.User
	handle := userService.userRepo.GetUserById(&user, id)
	if handle != nil {
		fmt.Println("Error")
	}
	return user, handle
}

func (userService *userService) GetUserByUserId(userId string) (users models.User, err error) {
	var user models.User
	handle := userService.userRepo.GetUserByUserId(&user, userId)
	if handle != nil {
		fmt.Println("Error")
	}
	return user, handle
}

func (userService *userService) InsertUser(user *models.User) (err error) {
	handle := userService.userRepo.InsertUser(user)
	if handle != nil {
		fmt.Println("Error")
	}
	return handle
}

func (userService *userService) UpdateUser(user *models.User, id string) (err error) {
	handle := userService.userRepo.UpdateUser(user)
	if handle != nil {
		fmt.Println("Error")
	}
	return handle
}

func (userService *userService) DeleteUserById(id string) (err error) {
	handle := userService.userRepo.DeleteUserById(id)
	if handle != nil {
		fmt.Println("Error")
	}
	return handle
}
