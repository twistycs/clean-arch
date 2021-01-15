package repositories

import (
	"github.com/twisty/clean-arch/imp"
	"github.com/twisty/clean-arch/models"
	"gorm.io/gorm"
)

type repository struct {
	connect *gorm.DB
}

func RepositoriesConstruct(connect *gorm.DB) imp.UserRepository {
	return &repository{connect}
}

func (repo *repository) GetAllUser(u *[]models.User) (err error) {
	if err = repo.connect.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repository) GetUserById(u *models.User, id string) (err error) {
	if err = repo.connect.Where("id = ?", id).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repository) GetUserByUserId(u *models.User, userId string) (err error) {
	if err = repo.connect.Where("user_id = ?", userId).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repository) InsertUser(u *models.User) (err error) {
	if err = repo.connect.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repository) UpdateUser(u *models.User) (err error) {
	if err = repo.connect.Save(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repository) DeleteUserById(id string) (err error) {
	if err = repo.connect.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
