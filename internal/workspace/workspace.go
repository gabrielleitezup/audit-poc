package workspace

import (
	"audit-poc/internal/workspace/models"
	"audit-poc/internal/workspace/payloads"
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

type Main struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) ServiceMethods {
	return Main{db}
}

func (main Main) ParseWorkspace(entity io.ReadCloser) (payloads.Request, error) {
	var newSubs *payloads.Request

	err := json.NewDecoder(entity).Decode(&newSubs)
	if err != nil {
		return payloads.Request{}, errors.New("Decode.Error")
	}

	return *newSubs, nil
}

func (main Main) SaveWorkspace(ctx context.Context, request payloads.Request) (payloads.Response, error) {
	entity := models.Workspace{
		Id:   uuid.New(),
		Name: request.Name,
	}

	res := main.db.WithContext(ctx).Model(&models.Workspace{}).Create(&entity)
	if res.Error != nil {
		logrus.Errorln("Save Workspace error:", res.Error)
		return payloads.Response{}, res.Error
	}

	return payloads.Response{Id: entity.Id, Name: entity.Name}, nil
}

func (main Main) UpdateWorkspace(ctx context.Context, request payloads.Request, workspaceId uuid.UUID) (payloads.Response, error) {
	var entity = models.Workspace{}

	res := main.db.WithContext(ctx).Model(&models.Workspace{}).First(&entity, workspaceId)
	if res.Error != nil {
		logrus.Errorln("Save Workspace error:", res.Error)
		return payloads.Response{}, res.Error
	}

	resUp := main.db.WithContext(ctx).Model(&entity).Update("name", request.Name).Scan(&entity)
	if resUp.Error != nil {
		logrus.Errorln("Save Workspace error:", res.Error)
		return payloads.Response{}, res.Error
	}

	return payloads.Response{Id: entity.Id, Name: entity.Name}, nil
}

func (main Main) DeleteWorkspace(ctx context.Context, workspaceId uuid.UUID) (payloads.Response, error) {
	var entity = models.Workspace{}

	res := main.db.WithContext(ctx).Model(&models.Workspace{}).First(&entity, workspaceId)
	if res.Error != nil {
		logrus.Errorln("Save Workspace error:", res.Error)
		return payloads.Response{}, res.Error
	}

	resDel := main.db.WithContext(ctx).Model(&models.Workspace{}).Delete(&entity)
	if resDel.Error != nil {
		logrus.Errorln("Save Workspace error:", resDel.Error)
		return payloads.Response{}, resDel.Error
	}

	return payloads.Response{Id: entity.Id, Name: entity.Name}, nil
}
