package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CircleUserGroup struct {
	Id          uuid.UUID      `json:"id"`
	CircleId    uuid.UUID      `json:"circle_id"`
	UserGroupId uuid.UUID      `json:"user_group_id"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}
