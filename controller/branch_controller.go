package controller

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/twisty/clean-arch/imp"
	"github.com/twisty/clean-arch/models"
)

var responseBranch models.Response

type BranchController struct {
	branchService imp.BranchService
}

func BranchControllerInit(branchService imp.BranchService) *BranchController {
	return &BranchController{
		branchService: branchService,
	}
}
func (u *BranchController) GetAllBranchController(c *gin.Context) {
	log.Info("Get All USER !!!")
	resp, err := u.branchService.GetAllBranch()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (u *BranchController) GetBranchByIdController(c *gin.Context) {
	id := c.Params.ByName("id")
	resp, err := u.branchService.GetBranchById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	} else if (models.Branch{} == resp) {
		c.JSON(http.StatusNotFound, "ID : "+id+" Not Found")
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (u *BranchController) InsertBranchController(c *gin.Context) {
	var jsonInputBranch models.Branch
	if err := c.ShouldBindJSON(&jsonInputBranch); err != nil {
		log.Fatal("Err : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := u.branchService.InsertBranch(&jsonInputBranch)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	} else {
		responseBranch.StatusCode = http.StatusCreated
		responseBranch.StatusMsg = "Created"
		c.JSON(http.StatusCreated, responseBranch)
	}
}

func (u *BranchController) UpdateBranchController(c *gin.Context) {
	var jsonInputBranch models.Branch
	id := c.Params.ByName("id")
	if err := c.ShouldBindJSON(&jsonInputBranch); err != nil {
		log.Error("Err : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		log.Info("[UPDATE CONTROLLER] Branch Input Information : ", jsonInputBranch)
		resp, err := u.branchService.GetBranchById(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		} else if (models.Branch{} == resp) {
			responseBranch.StatusCode = http.StatusNotFound
			responseBranch.StatusMsg = "ID : " + id + " Not Found"
			c.JSON(http.StatusNotFound, responseBranch)
		} else {
			if err := u.branchService.UpdateBranch(&jsonInputBranch, id); err != nil {
				log.Error("Err : ", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			responseBranch.StatusCode = http.StatusOK
			responseBranch.StatusMsg = "Updated"
			c.JSON(http.StatusOK, responseBranch)
		}
	}
}

func (u *BranchController) DeleteBranchByIdController(c *gin.Context) {
	id := c.Params.ByName("id")
	log.Info("[Branch Controller] Delete Branch ID : ", id)
	resp, err := u.branchService.GetBranchById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	} else if (models.Branch{} == resp) {
		responseBranch.StatusCode = http.StatusNotFound
		responseBranch.StatusMsg = "ID : " + id + " Not Found"
		c.JSON(http.StatusNotFound, responseBranch)
	} else {
		if err := u.branchService.DeleteBranchById(id); err != nil {
			log.Error("Err : ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		responseBranch.StatusCode = http.StatusOK
		responseBranch.StatusMsg = "Deleted"
		c.JSON(http.StatusOK, responseBranch)
	}
}
