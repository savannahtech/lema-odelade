package entity

import (
	"gorm.io/gorm"
)

type Commit struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey"`
	Message  string
	Author   string
	Date     string
	Url      string
	Sha      string
	RepoName string
}
