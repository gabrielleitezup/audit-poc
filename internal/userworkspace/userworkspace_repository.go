package userworkspace

import (
	"audit-poc/internal/userworkspace/models"
	"audit-poc/internal/userworkspace/payloads"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
)

type ServiceMethods interface {
	ParseUserWorkspace(entity io.ReadCloser) (payloads.Request, error)
	AssociateUserGroupToWorkspace(ctx context.Context, request payloads.Request, workspaceId uuid.UUID) (payloads.Response, error)
}

type UserGroupWorkspaceRepository struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) ServiceMethods {
	return UserGroupWorkspaceRepository{db}
}

func (main UserGroupWorkspaceRepository) ParseUserWorkspace(entity io.ReadCloser) (payloads.Request, error) {
	var newSubs *payloads.Request

	err := json.NewDecoder(entity).Decode(&newSubs)
	if err != nil {
		return payloads.Request{}, errors.New("Decode.Error")
	}

	return *newSubs, nil
}

func (main UserGroupWorkspaceRepository) AssociateUserGroupToWorkspace(ctx context.Context, request payloads.Request, workspaceId uuid.UUID) (payloads.Response, error) {
	entity := models.UserGroupWorkspace{
		Id:          uuid.New(),
		UserGroupId: request.GroupId,
		WorkspaceId: workspaceId,
		Permission:  request.Permissions,
	}

	err := main.AssociateWithContext(ctx, entity)
	if err != nil {
		logrus.Errorln("Save UserGroup error:", err)
		return payloads.Response{}, err
	}

	return payloads.Response{Id: entity.Id, GroupId: entity.UserGroupId, WorkspaceId: entity.WorkspaceId}, nil
}
