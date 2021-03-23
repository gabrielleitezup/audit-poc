package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Circle struct {
	Id          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Rules       []byte         `json:"rules" gorm:"type:jsonb"`
	WorkspaceId uuid.UUID      `json:"workspace_id"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}
