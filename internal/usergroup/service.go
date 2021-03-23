package usergroup

import (
	"audit-poc/internal/usergroup/models"
	"audit-poc/internal/usergroup/payloads"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
)

type ServiceMethods interface {
	ParseUserGroup(entity io.ReadCloser) (payloads.Request, error)
	SaveUserGroup(ctx context.Context, request payloads.Request) (payloads.Response, error)
	UpdateUserGroup(ctx context.Context, request payloads.Request, workspaceId uuid.UUID) (payloads.Response, error)
	DeleteUserGroup(ctx context.Context, workspaceId uuid.UUID) (payloads.Response, error)
}

type Main struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) ServiceMethods {
	return Main{db}
}


func (main Main) ParseUserGroup(entity io.ReadCloser) (payloads.Request, error) {
	var newSubs *payloads.Request

	err := json.NewDecoder(entity).Decode(&newSubs)
	if err != nil {
		return payloads.Request{}, errors.New("Decode.Error")
	}

	return *newSubs, nil
}

func (main Main) SaveUserGroup(ctx context.Context, request payloads.Request) (payloads.Response, error) {
	entity := models.UserGroup{
		Id:   uuid.New(),
		Name: request.Name,
	}

	res := main.db.WithContext(ctx).Model(&models.UserGroup{}).Create(&entity)
	if res.Error != nil {
		logrus.Errorln("Save UserGroup error:", res.Error)
		return payloads.Response{}, res.Error
	}

	return payloads.Response{Id: entity.Id, Name: entity.Name}, nil
}

func (main Main) UpdateUserGroup(ctx context.Context, request payloads.Request, workspaceId uuid.UUID) (payloads.Response, error) {
	var entity = models.UserGroup{}

	res := main.db.WithContext(ctx).Model(&models.UserGroup{}).First(&entity, workspaceId)
	if res.Error != nil {
		logrus.Errorln("Update UserGroup error:", res.Error)
		return payloads.Response{}, res.Error
	}

	resUp := main.db.WithContext(ctx).Model(&entity).Update("name", request.Name).Scan(&entity)
	if resUp.Error != nil {
		logrus.Errorln("Update UserGroup error:", res.Error)
		return payloads.Response{}, res.Error
	}

	return payloads.Response{Id: entity.Id, Name: entity.Name}, nil
}

func (main Main) DeleteUserGroup(ctx context.Context, workspaceId uuid.UUID) (payloads.Response, error) {
	var entity = models.UserGroup{}

	res := main.db.WithContext(ctx).Model(&models.UserGroup{}).First(&entity, workspaceId)
	if res.Error != nil {
		logrus.Errorln("Delete UserGroup error:", res.Error)
		return payloads.Response{}, res.Error
	}

	resDel := main.db.WithContext(ctx).Model(&models.UserGroup{}).Delete(&entity)
	if resDel.Error != nil {
		logrus.Errorln("Delete UserGroup error:", resDel.Error)
		return payloads.Response{}, resDel.Error
	}

	return payloads.Response{Id: entity.Id, Name: entity.Name}, nil
}

