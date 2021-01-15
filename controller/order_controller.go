package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/twisty/clean-arch/imp"
	"github.com/twisty/clean-arch/models"
)

type OrderController struct {
	orderService imp.OrderService
	userService  imp.UserService
}

func OrderControllerInit(orderService imp.OrderService, userService imp.UserService) *OrderController {
	return &OrderController{
		orderService: orderService,
		userService:  userService,
	}
}
func (o *OrderController) GetAllOrderController(c *gin.Context) {
	resp, err := o.orderService.GetAllOrder()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}
func (o *OrderController) InsertOrderController(c *gin.Context) {
	var jsonInputOrder models.Order
	if err := c.ShouldBindJSON(&jsonInputOrder); err != nil {
		log.Error("Err : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Info(jsonInputOrder)
	user, err := o.userService.GetUserByUserId(jsonInputOrder.User.UserID)
	jsonInputOrder.User = user
	log.Info(jsonInputOrder)
	if err != nil {
		log.Error(err.Error())
	} else {
		c.JSON(http.StatusOK, jsonInputOrder)
	}

}
