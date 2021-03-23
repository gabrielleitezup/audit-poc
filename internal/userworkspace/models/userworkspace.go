package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGroupWorkspace struct {
	Id          uuid.UUID      `json:"id"`
	UserGroupId uuid.UUID      `json:"user_group_id"`
	WorkspaceId uuid.UUID      `json:"workspace_id"`
	Permission  []byte         `json:"permission" gorm:"type:jsonb"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}