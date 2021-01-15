package imp

import "github.com/twisty/clean-arch/models"

type BranchService interface {
	GetAllBranch() (userList []models.Branch, err error)
	GetBranchById(id string) (user models.Branch, err error)
	GetBranchByBranchName(name string) (user models.Branch, err error)
	InsertBranch(u *models.Branch) (err error)
	UpdateBranch(u *models.Branch, id string) (err error)
	DeleteBranchById(id string) (err error)
}

type BranchRepository interface {
	GetAllBranch(u *[]models.Branch) (err error)
	GetBranchById(u *models.Branch, id string) (err error)
	GetBranchByBranchName(u *models.Branch, name string) (err error)
	InsertBranch(u *models.Branch) (err error)
	UpdateBranch(u *models.Branch) (err error)
	DeleteBranchById(id string) (err error)
}
