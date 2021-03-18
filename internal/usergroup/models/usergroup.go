package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGroup struct {
	Id        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	DeletedAt gorm.DeletedAt `json:"-"`
}