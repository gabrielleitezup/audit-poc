package workspace

import (
	"audit-poc/internal/userworkspace/workspace/models"
	"audit-poc/internal/userworkspace/workspace/payloads"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
)

type ServiceMethods interface {
	SaveWorkspace(ctx context.Context, request payloads.Request) (payloads.Response, error)
	UpdateWorkspace(ctx context.Context, request payloads.Request, workspaceId uuid.UUID) (payloads.Response, error)
	DeleteWorkspace(ctx context.Context, workspaceId uuid.UUID) (payloads.Response, error)
	ParseWorkspace(subscription io.ReadCloser) (payloads.Request, error)
}

type WorkspaceRepository struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) ServiceMethods {
	return WorkspaceRepository{db}
}

func (main WorkspaceRepository) ParseWorkspace(entity io.ReadCloser) (payloads.Request, error) {
	var newSubs *payloads.Request

	err := json.NewDecoder(entity).Decode(&newSubs)
	if err != nil {
		return payloads.Request{}, errors.New("Decode.Error")
	}

	return *newSubs, nil
}

func (main WorkspaceRepository) SaveWorkspace(ctx context.Context, request payloads.Request) (payloads.Response, error) {
	entity := models.Workspace{
		Id:   uuid.New(),
		Name: request.Name,
	}

	err := main.SaveWithContext(ctx, entity)
	if err != nil {
		logrus.Errorln("Save Workspace error:", err)
		return payloads.Response{}, err
	}

	return payloads.Response{Id: entity.Id, Name: entity.Name}, nil
}

func (main WorkspaceRepository) UpdateWorkspace(ctx context.Context, request payloads.Request, workspaceId uuid.UUID) (payloads.Response, error) {
	var entity = models.Workspace{}

	res, err := main.UpdateWithContext(ctx, entity, request, workspaceId)
	if err != nil {
		logrus.Errorln("Update Workspace error:", err)
		return payloads.Response{}, err
	}

	return payloads.Response{Id: res.Id, Name: res.Name}, nil
}

func (main WorkspaceRepository) DeleteWorkspace(ctx context.Context, workspaceId uuid.UUID) (payloads.Response, error) {
	var entity = models.Workspace{}

	err := main.DeleteWithContext(ctx, entity, workspaceId)
	if err != nil {
		logrus.Errorln("Save Workspace error:", err)
		return payloads.Response{}, err
	}

	return payloads.Response{}, nil
}
