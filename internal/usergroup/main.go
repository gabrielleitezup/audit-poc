package usergroup

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
)

type ServiceMethods interface {
	ParseUserGroup(entity io.ReadCloser) (Request, error)
	SaveUserGroup(ctx context.Context, request Request) (Response, error)
	UpdateUserGroup(ctx context.Context, request Request, workspaceId uuid.UUID) (Response, error)
	DeleteUserGroup(ctx context.Context, workspaceId uuid.UUID) (Response, error)
}

type Main struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) ServiceMethods {
	return Main{db}
}
