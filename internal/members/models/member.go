package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	Id          uuid.UUID      `json:"id"`
	UserGroupId uuid.UUID      `json:"user_group_id"`
	Username    string         `json:"username"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}
