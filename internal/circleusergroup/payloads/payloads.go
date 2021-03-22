package payloads

import (
	"github.com/google/uuid"
)

type Request struct {
	UserGroupId uuid.UUID `json:"userGroupId"`
}

type Response struct {
	CircleId    uuid.UUID `json:"circleId"`
	UserGroupId uuid.UUID `json:"userGroupId"`
}
