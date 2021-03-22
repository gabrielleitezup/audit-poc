package payloads

import "github.com/google/uuid"

type Request struct {
	Username string `json:"username"`
}

type Response struct {
	Id          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	UserGroupId uuid.UUID `json:"userGroupId"`
}
