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
func (o *OrderController) InsertOrderController(c *gin.Context) {
	var jsonInputOrder models.Order
	if err := c.ShouldBindJSON(&jsonInputOrder); err != nil {
		log.Error("Err : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := o.orderService.InsertOrder(&jsonInputOrder)
	if err.Error() == "UserID" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Id Not Found"})
	} else if err != nil {
		switch errorCase := err.Error(); {
		case errorCase == "UserID not found":
			c.JSON(http.StatusNotFound, gin.H{"error": "User Id Not Found"})
			break
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, jsonInputOrder)
	}

}
