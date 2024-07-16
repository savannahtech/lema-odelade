package entity

import (
	"gorm.io/gorm"
)

type RepoDetail struct {
	gorm.Model
	ID             uint64 `gorm:"primaryKey"`
	Name           string
	Owner          string
	Description    string
	Url            string
	Language       string
	ForkCount      int
	StarCount      int
	OpenIssueCount int
	WatcherCount   int
	DateCreated    string
	UpdatedDate    string
}
