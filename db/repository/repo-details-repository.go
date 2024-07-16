package repository

import (
	Entity "accessment.com/microservice/db/entity"
	"accessment.com/microservice/db/postgres"
	"gorm.io/gorm"
)

type RepoDetailRepository interface {
	Store(repoDetail Entity.RepoDetail) error
	GetByName(name string) (Entity.RepoDetail, error)
	GetAll() ([]Entity.RepoDetail, error)
}

type RepoDetailRepo struct {
}

func (rep *RepoDetailRepo) Store(repoDetail Entity.RepoDetail) error {
	var database *gorm.DB = postgres.ConnectToDb()
	err := database.Create(&repoDetail).Error
	return err
}

func (rep *RepoDetailRepo) GetByName(name string) (Entity.RepoDetail, error) {
	var database *gorm.DB = postgres.ConnectToDb()
	var repoDetail Entity.RepoDetail
	err := database.Where(&Entity.RepoDetail{Name: name}).Find(&repoDetail).Error
	return repoDetail, err
}

func (rep *RepoDetailRepo) GetAll() ([]Entity.RepoDetail, error) {
	var database *gorm.DB = postgres.ConnectToDb()
	var repoDetail []Entity.RepoDetail
	err := database.Find(&repoDetail).Error
	return repoDetail, err
}
