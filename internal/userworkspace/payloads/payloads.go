package payloads

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Request struct {
	GroupId     uuid.UUID       `json:"groupId"`
	Permissions json.RawMessage `json:"permissions"`
}

type Response struct {
	Id          uuid.UUID `json:"id"`
	GroupId     uuid.UUID `json:"groupId"`
	WorkspaceId uuid.UUID `json:"workspaceId"`
}
