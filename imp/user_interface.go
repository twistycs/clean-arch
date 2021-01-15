package imp

import "github.com/twisty/clean-arch/models"

type UserService interface {
	GetAllUser() (userList []models.User, err error)
	GetUserById(id string) (user models.User, err error)
	GetUserByUserId(userId string) (user models.User, err error)
	InsertUser(u *models.User) (err error)
	UpdateUser(u *models.User, id string) (err error)
	DeleteUserById(id string) (err error)
}

type UserRepository interface {
	GetAllUser(u *[]models.User) (err error)
	GetUserById(u *models.User, id string) (err error)
	GetUserByUserId(u *models.User, userId string) (err error)
	InsertUser(u *models.User) (err error)
	UpdateUser(u *models.User) (err error)
	DeleteUserById(id string) (err error)
}
