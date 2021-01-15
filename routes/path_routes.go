package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/twisty/clean-arch/controller"
	"github.com/twisty/clean-arch/database"
	"github.com/twisty/clean-arch/repositories"
	"github.com/twisty/clean-arch/services"
)

func SetUpRoutes() *gin.Engine {

	userRepo := repositories.RepositoriesConstruct(database.DB)
	userService := services.NewService(userRepo)
	userController := controller.UserControllerInit(userService)

	r := gin.Default()
	user := r.Group("/v1/users")
	{
		user.GET("/", userController.GetAllUserController)
	}

	{
		user.GET("/:id", userController.GetUserByIdController)
	}

	{
		user.POST("/user", userController.InsertUserController)
	}

	orderRepo := repositories.RepositoriesOrder(database.DB)
	orderService := services.NewOrderService(orderRepo)
	orderController := controller.OrderControllerInit(orderService)

	order := r.Group("/v1/order")
	{
		order.GET("/", orderController.GetAllOrderController)
	}

	return r
}
