package models

import (
	"github.com/google/uuid"
	"time"
)

type Audition struct {
	Id           uuid.UUID  `json:"id"`
	Username     string     `json:"username"`
	TableName    string     `json:"table_name"`
	Operation    string     `json:"operation"`
	EntityId     string     `json:"entity_id"`
	CurrentState []byte     `json:"current_state" gorm:"type:jsonb"`
	UserIpAddr   string     `json:"user_ip_addr"`
	UserAgent    string     `json:"user_agent"`
	CreatedAt    *time.Time `json:"createdAt"`
}
