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
