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

type Main struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) ServiceMethods {
	return Main{db}
}

func (main Main) ParseUserWorkspace(entity io.ReadCloser) (payloads.Request, error) {
	var newSubs *payloads.Request

	err := json.NewDecoder(entity).Decode(&newSubs)
	if err != nil {
		return payloads.Request{}, errors.New("Decode.Error")
	}

	return *newSubs, nil
}

func (main Main) AssociateUserGroupToWorkspace(ctx context.Context, request payloads.Request, workspaceId uuid.UUID) (payloads.Response, error) {
	entity := models.UserGroupWorkspace{
		Id:          uuid.New(),
		UserGroupId: request.GroupId,
		WorkspaceId: workspaceId,
		Permission:  request.Permissions,
	}

	res := main.db.WithContext(ctx).Model(&models.UserGroupWorkspace{}).Create(&entity)
	if res.Error != nil {
		logrus.Errorln("Save Workspace error:", res.Error)
		return payloads.Response{}, res.Error
	}

	return payloads.Response{Id: entity.Id, GroupId: entity.UserGroupId, WorkspaceId: entity.WorkspaceId}, nil
}
