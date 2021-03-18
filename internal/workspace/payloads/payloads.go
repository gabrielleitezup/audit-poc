package payloads

import "github.com/google/uuid"

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}