package payloads

import (
	"github.com/google/uuid"
)

type Request struct {
	Name        string          `json:"name"`
	CircleId    uuid.UUID      `json:"circleId"`
	UserGroupId uuid.UUID      `json:"userGroupId"`
}

type Response struct {
	CircleId    uuid.UUID `json:"circleId"`
	UserGroupId uuid.UUID `json:"userGroupId"`
}
