package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Workspace struct {
	Id         uuid.UUID      `json:"id"`
	Name       string         `json:"name"`
	MatcherUrl string         `json:"matcher_url"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}
