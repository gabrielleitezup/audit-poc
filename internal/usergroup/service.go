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

type UserGroupRepository struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) ServiceMethods {
	return UserGroupRepository{db}
}

func (main UserGroupRepository) ParseUserGroup(entity io.ReadCloser) (payloads.Request, error) {
	var newSubs *payloads.Request

	err := json.NewDecoder(entity).Decode(&newSubs)
	if err != nil {
		return payloads.Request{}, errors.New("Decode.Error")
	}

	return *newSubs, nil
}

func (main UserGroupRepository) SaveUserGroup(ctx context.Context, request payloads.Request) (payloads.Response, error) {
	entity := models.UserGroup{
		Id:   uuid.New(),
		Name: request.Name,
	}

	err := main.SaveWithContext(ctx, entity)
	if err != nil {
		logrus.Errorln("Save UserGroup error:", err)
		return payloads.Response{}, err
	}

	return payloads.Response{Id: entity.Id, Name: entity.Name}, nil
}

func (main UserGroupRepository) UpdateUserGroup(ctx context.Context, request payloads.Request, userGroupId uuid.UUID) (payloads.Response, error) {
	var entity = models.UserGroup{}

	res, err := main.UpdateWithContext(ctx, entity, request, userGroupId)
	if err != nil {
		logrus.Errorln("Update UserGroup error:", err)
		return payloads.Response{}, err
	}

	return payloads.Response{Id: res.Id, Name: res.Name}, nil
}

func (main UserGroupRepository) DeleteUserGroup(ctx context.Context, userGroupId uuid.UUID) (payloads.Response, error) {
	var entity = models.UserGroup{}

	err := main.DeleteWithContext(ctx, entity, userGroupId)
	if err != nil {
		logrus.Errorln("Delete UserGroup error:", err)
		return payloads.Response{}, err
	}

	return payloads.Response{Id: entity.Id, Name: entity.Name}, nil
}

