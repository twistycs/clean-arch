package services

import (
	"fmt"

	"github.com/twisty/clean-arch/imp"
	"github.com/twisty/clean-arch/models"
)

type orderService struct {
	orderRepo imp.OrderRepository
}

func NewOrderService(repo imp.OrderRepository) imp.OrderService {
	return &orderService{orderRepo: repo}
}

func (orderService *orderService) GetAllOrder() (ordes []models.Order, err error) {
	var order []models.Order
	handle := orderService.orderRepo.GetAllOrder(&order)
	if handle != nil {
		fmt.Println("Error")
	}
	return order, handle
}

// func (orderService *orderService) InsertOrder(user *models.User) (err error) {
// 	t := time.Now()
// 	fmt.Println(t.String())

// 	orderNo, err := strconv.ParseInt(strings.Replace(t.Format("20060102150405.0000"), ".", "", -1), 10, 64)
// 	if err != nil {
// 		log.Error(err.Error())
// 	}
// 	fmt.Println(orderNo)
// }
