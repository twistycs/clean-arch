package repositories

import (
	"github.com/twisty/clean-arch/imp"
	"github.com/twisty/clean-arch/models"
	"gorm.io/gorm"
)

type repositoryBranch struct {
	connect *gorm.DB
}

func RepositoriesBranch(connect *gorm.DB) imp.BranchRepository {
	return &repositoryBranch{connect}
}

func (repo *repositoryBranch) GetAllBranch(b *[]models.Branch) (err error) {
	if err = repo.connect.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repositoryBranch) GetBranchById(u *models.Branch, id string) (err error) {
	if err = repo.connect.Where("id = ?", id).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repositoryBranch) GetBranchByBranchName(u *models.Branch, name string) (err error) {
	if err = repo.connect.Where("name = ?", name).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repositoryBranch) InsertBranch(u *models.Branch) (err error) {
	if err = repo.connect.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repositoryBranch) UpdateBranch(u *models.Branch) (err error) {
	if err = repo.connect.Save(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repositoryBranch) DeleteBranchById(id string) (err error) {
	if err = repo.connect.Delete(&models.Branch{}, id).Error; err != nil {
		return err
	}
	return nil
}
