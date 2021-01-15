package repositories

import (
	"github.com/twisty/clean-arch/imp"
	"github.com/twisty/clean-arch/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderRepository struct {
	connect *gorm.DB
}

func RepositoriesOrder(connect *gorm.DB) imp.OrderRepository {
	return &orderRepository{connect}
}

func (repo *orderRepository) GetAllOrder(o *[]models.Order) (err error) {
	// if err = repo.connect.Preload(clause.Associations).Find(o).Error; err != nil {
	if err = repo.connect.Preload(clause.Associations).Find(o).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repository) InsertOrder(u *models.Order) (err error) {
	if err = repo.connect.Create(u).Error; err != nil {
		return err
	}
	return nil
}
