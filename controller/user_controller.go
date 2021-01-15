package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/twisty/clean-arch/imp"
	"github.com/twisty/clean-arch/models"
)

var response models.Response

type UserController struct {
	userService imp.UserService
}

func UserControllerInit(userService imp.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
func (u *UserController) GetAllUserController(c *gin.Context) {
	resp, err := u.userService.GetAllUser()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (u *UserController) GetUserByIdController(c *gin.Context) {
	id := c.Params.ByName("id")
	resp, err := u.userService.GetUserById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	} else if len(resp) < 1 {
		c.JSON(http.StatusNotFound, "ID : "+id+" Not Found")
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (u *UserController) InsertUserController(c *gin.Context) {
	var jsonInputUser models.User
	if err := c.ShouldBindJSON(&jsonInputUser); err != nil {
		log.Fatal("Err : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := u.userService.InsertUser(&jsonInputUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	} else {
		response.StatusCode = http.StatusCreated
		response.StatusMsg = "Created"
		c.JSON(http.StatusCreated, response)
	}
}
