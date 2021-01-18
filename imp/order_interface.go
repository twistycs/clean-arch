package imp

import "github.com/twisty/clean-arch/models"

type OrderService interface {
	GetAllOrder() (orderList []models.Order, err error)
	InsertOrder(o *models.Order) (err error)
}

type OrderRepository interface {
	GetAllOrder(o *[]models.Order) (err error)
	InsertOrder(o *models.Order) (err error)
}
