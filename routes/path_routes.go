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

	branchRepo := repositories.RepositoriesBranch(database.DB)
	branchService := services.NewBranchService(branchRepo)
	branchController := controller.BranchControllerInit(branchService)

	orderRepo := repositories.RepositoriesOrder(database.DB)
	orderService := services.NewOrderService(orderRepo, userRepo, branchRepo)
	orderController := controller.OrderControllerInit(orderService)

	r := gin.Default()
	user := r.Group("/v1/users")
	{
		user.GET("/", userController.GetAllUserController)
	}

	{
		user.GET("/:id", userController.GetUserByIdController)
	}

	{
		user.POST("/", userController.InsertUserController)
	}

	{
		user.PUT("/:id", userController.UpdateUserController)
	}

	{
		user.DELETE("/:id", userController.DeleteUserByIdController)
	}

	order := r.Group("/v1/orders")
	{
		order.GET("/", orderController.GetAllOrderController)
	}
	{
		order.POST("/", orderController.InsertOrderController)
	}

	branch := r.Group("/v1/branches")
	{
		branch.GET("/", branchController.GetAllBranchController)
	}

	{
		branch.GET("/:id", branchController.GetBranchByIdController)
	}

	{
		branch.POST("/", branchController.InsertBranchController)
	}

	{
		branch.PUT("/:id", branchController.UpdateBranchController)
	}

	{
		branch.DELETE("/:id", branchController.DeleteBranchByIdController)
	}

	return r
}
