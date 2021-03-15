package userworkspace

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
)

type ServiceMethods interface {
	ParseUserWorkspace(entity io.ReadCloser) (Request, error)
	AssociateUserGroupToWorkspace(ctx context.Context, request Request, workspaceId uuid.UUID) (Response, error)
}

type Main struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) ServiceMethods {
	return Main{db}
}
