package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Deployment struct {
	Id        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Version   string         `json:"version"`
	CircleId  uuid.UUID      `json:"circle_id"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
