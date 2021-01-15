package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/twisty/clean-arch/imp"
)

type OrderController struct {
	orderService imp.OrderService
}

func OrderControllerInit(orderService imp.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
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
