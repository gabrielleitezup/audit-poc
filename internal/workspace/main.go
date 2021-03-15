package workspace

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
)

type ServiceMethods interface {
	SaveWorkspace(ctx context.Context, request Request) (Response, error)
	UpdateWorkspace(ctx context.Context, request Request, workspaceId uuid.UUID) (Response, error)
	DeleteWorkspace(ctx context.Context, workspaceId uuid.UUID) (Response, error)
	ParseWorkspace(subscription io.ReadCloser) (Request, error)
}

type Main struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) ServiceMethods {
	return Main{db}
}
