package payloads

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Request struct {
	Name  string          `json:"name"`
	WorkspaceId uuid.UUID `json:"workspaceId"`
	Rules json.RawMessage `json:"rules"`
}

type Response struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
