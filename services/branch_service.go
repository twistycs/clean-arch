package services

import (
	"fmt"

	"github.com/twisty/clean-arch/imp"
	"github.com/twisty/clean-arch/models"
)

type branchService struct {
	branchRepo imp.BranchRepository
}

func NewBranchService(repo imp.BranchRepository) imp.BranchService {
	return &branchService{branchRepo: repo}
}

func (branchService *branchService) GetAllBranch() (branchs []models.Branch, err error) {
	var branch []models.Branch
	handle := branchService.branchRepo.GetAllBranch(&branch)
	if handle != nil {
		fmt.Println("Error")
	}
	return branch, handle
}

func (branchService *branchService) GetBranchById(id string) (branchs models.Branch, err error) {
	var branch models.Branch
	handle := branchService.branchRepo.GetBranchById(&branch, id)
	if handle != nil {
		fmt.Println("Error")
	}
	return branch, handle
}

func (branchService *branchService) GetBranchByBranchName(name string) (branchs models.Branch, err error) {
	var branch models.Branch
	handle := branchService.branchRepo.GetBranchByBranchName(&branch, name)
	if handle != nil {
		fmt.Println("Error")
	}
	return branch, handle
}

func (branchService *branchService) InsertBranch(branch *models.Branch) (err error) {
	handle := branchService.branchRepo.InsertBranch(branch)
	if handle != nil {
		fmt.Println("Error")
	}
	return handle
}

func (branchService *branchService) UpdateBranch(branch *models.Branch, id string) (err error) {
	handle := branchService.branchRepo.UpdateBranch(branch)
	if handle != nil {
		fmt.Println("Error")
	}
	return handle
}

func (branchService *branchService) DeleteBranchById(id string) (err error) {
	handle := branchService.branchRepo.DeleteBranchById(id)
	if handle != nil {
		fmt.Println("Error")
	}
	return handle
}
