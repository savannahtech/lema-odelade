package repository

import (
	Entity "accessment.com/microservice/db/entity"
	"accessment.com/microservice/db/postgres"
	"gorm.io/gorm"
)

type CommitRepository interface {
	Store(commit Entity.Commit) error
	GetCommit(repoName string) ([]Entity.Commit, error)
	GetCommitInSha(shaList []string) ([]Entity.Commit, error)
	StoreList(commits []Entity.Commit) error
}

type CommitRepo struct {
}

func (com *CommitRepo) Store(commit Entity.Commit) error {
	var database *gorm.DB = postgres.ConnectToDb()
	err := database.Create(&commit).Error
	return err
}

func (com *CommitRepo) StoreList(commits []Entity.Commit) error {
	var database *gorm.DB = postgres.ConnectToDb()
	err := database.Create(&commits).Error
	return err
}

func (com *CommitRepo) GetCommit(repoName string) ([]Entity.Commit, error) {
	var database *gorm.DB = postgres.ConnectToDb()
	var commits []Entity.Commit
	err := database.Where(&Entity.Commit{RepoName: repoName}).Find(&commits).Error
	return commits, err
}

func (com *CommitRepo) GetCommitInSha(shaList []string) ([]Entity.Commit, error) {
	var database *gorm.DB = postgres.ConnectToDb()
	var commits []Entity.Commit
	err := database.Where("commits.sha IN ?", shaList).Find(&commits).Error
	return commits, err
}
