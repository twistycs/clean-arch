package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/twisty/clean-arch/imp"
	"github.com/twisty/clean-arch/models"
)

type orderService struct {
	orderRepo  imp.OrderRepository
	userRepo   imp.UserRepository
	branchRepo imp.BranchRepository
}

func NewOrderService(repo imp.OrderRepository, userRepo imp.UserRepository, branchRepo imp.BranchRepository) imp.OrderService {
	return &orderService{orderRepo: repo, userRepo: userRepo, branchRepo: branchRepo}
}

func (orderService *orderService) GetAllOrder() (ordes []models.Order, err error) {
	var order []models.Order
	handle := orderService.orderRepo.GetAllOrder(&order)
	if handle != nil {
		fmt.Println("Error")
	}
	return order, handle
}

func (orderService *orderService) InsertOrder(jsonInputOrder *models.Order) (err error) {
	t := time.Now()
	fmt.Println(t.String())
	orderNo, err := strconv.ParseInt(strings.Replace(t.Format("20060102150405.0000"), ".", "", -1), 10, 64)
	jsonInputOrder.OrderNo = uint(orderNo)
	jsonInputOrder.CreatedDate = t
	jsonInputOrder.UpdatedDate = t
	var checkUser models.User
	var checkBranch models.Branch
	err = orderService.userRepo.GetUserByUserId(&checkUser, jsonInputOrder.User.UserID)
	if (models.User{} == checkUser) {
		return errors.New("UserID not found")
	}
	err = orderService.branchRepo.GetBranchByBranchName(&checkBranch, jsonInputOrder.Branch.Name)
	err = orderService.orderRepo.InsertOrder(jsonInputOrder)
	if err != nil {
		log.Error(err.Error())
	}
	return err
}
