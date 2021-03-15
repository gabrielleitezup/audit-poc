package usergroup

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
)

type UserGroup struct {
	Id        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (main Main) ParseUserGroup(entity io.ReadCloser) (Request, error) {
	var newSubs *Request

	err := json.NewDecoder(entity).Decode(&newSubs)
	if err != nil {
		return Request{}, errors.New("Decode.Error")
	}

	return *newSubs, nil
}

func (main Main) SaveUserGroup(ctx context.Context, request Request) (Response, error) {
	entity := UserGroup{
		Id:   uuid.New(),
		Name: request.Name,
	}

	res := main.db.WithContext(ctx).Model(&UserGroup{}).Create(&entity)
	if res.Error != nil {
		logrus.Errorln("Save Workspace error:", res.Error)
		return Response{}, res.Error
	}

	return Response{Id: entity.Id, Name: entity.Name}, nil
}

func (main Main) UpdateUserGroup(ctx context.Context, request Request, workspaceId uuid.UUID) (Response, error) {
	var entity = UserGroup{}

	res := main.db.WithContext(ctx).Model(&UserGroup{}).First(&entity, workspaceId)
	if res.Error != nil {
		logrus.Errorln("Update UserGroup error:", res.Error)
		return Response{}, res.Error
	}

	resUp := main.db.WithContext(ctx).Model(&entity).Update("name", request.Name).Scan(&entity)
	if resUp.Error != nil {
		logrus.Errorln("Update UserGroup error:", res.Error)
		return Response{}, res.Error
	}

	return Response{Id: entity.Id, Name: entity.Name}, nil
}

func (main Main) DeleteUserGroup(ctx context.Context, workspaceId uuid.UUID) (Response, error) {
	var entity = UserGroup{}

	res := main.db.WithContext(ctx).Model(&UserGroup{}).First(&entity, workspaceId)
	if res.Error != nil {
		logrus.Errorln("Delete UserGroup error:", res.Error)
		return Response{}, res.Error
	}

	resDel := main.db.WithContext(ctx).Model(&UserGroup{}).Delete(&entity)
	if resDel.Error != nil {
		logrus.Errorln("Delete UserGroup error:", resDel.Error)
		return Response{}, resDel.Error
	}

	return Response{Id: entity.Id, Name: entity.Name}, nil
}

