package payloads

import (
	"github.com/google/uuid"
	"time"
)

type Response struct {
	Id           uuid.UUID  `json:"id"`
	Username     string     `json:"username"`
	TableName    string     `json:"tableName"`
	Operation    string     `json:"operation"`
	EntityId     string     `json:"entityId"`
	CurrentState []byte     `json:"currentState"`
	UserIpAddr   string     `json:"userIpAddr"`
	UserAgent    string     `json:"userAgent"`
	CreatedAt    *time.Time `json:"createdAt"`
}