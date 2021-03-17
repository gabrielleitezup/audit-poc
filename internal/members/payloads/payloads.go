package payloads

import "github.com/google/uuid"

type Request struct {
	Username string `json:"username"`
}

type Response struct {
	Id string `json:"id"`
	Username string `json:"username"`
	UserGroupId uuid.UUID `json:"user_group_id"`
}