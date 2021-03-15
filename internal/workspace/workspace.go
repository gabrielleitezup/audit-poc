package workspace

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
)

type Workspace struct {
	Id         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	MatcherUrl string     `json:"matcher_url"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}

func (main Main) ParseWorkspace(entity io.ReadCloser) (Request, error) {
	var newSubs *Request

	err := json.NewDecoder(entity).Decode(&newSubs)
	if err != nil {
		return Request{}, errors.New("Decode.Error")
	}

	return *newSubs, nil
}

func (main Main) SaveWorkspace(ctx context.Context, request Request) (Response, error) {
	entity := Workspace{
		Id:   uuid.New(),
		Name: request.Name,
	}

	res := main.db.WithContext(ctx).Model(&Workspace{}).Create(&entity)
	if res.Error != nil {
		logrus.Errorln("Save Workspace error:", res.Error)
		return Response{}, res.Error
	}

	return Response{Id: entity.Id, Name: entity.Name}, nil
}

func (main Main) UpdateWorkspace(ctx context.Context, request Request, workspaceId uuid.UUID) (Response, error) {
	var entity = Workspace{}

	res := main.db.WithContext(ctx).Model(&Workspace{}).First(&entity, workspaceId)
	if res.Error != nil {
		logrus.Errorln("Save Workspace error:", res.Error)
		return Response{}, res.Error
	}

	resUp := main.db.WithContext(ctx).Model(&entity).Update("name", request.Name).Scan(&entity)
	if resUp.Error != nil {
		logrus.Errorln("Save Workspace error:", res.Error)
		return Response{}, res.Error
	}

	return Response{Id: entity.Id, Name: entity.Name}, nil
}

func (main Main) DeleteWorkspace(ctx context.Context, workspaceId uuid.UUID) (Response, error) {
	var entity = Workspace{}

	res := main.db.WithContext(ctx).Model(&Workspace{}).First(&entity, workspaceId)
	if res.Error != nil {
		logrus.Errorln("Save Workspace error:", res.Error)
		return Response{}, res.Error
	}

	resDel := main.db.WithContext(ctx).Model(&Workspace{}).Delete(&entity)
	if resDel.Error != nil {
		logrus.Errorln("Save Workspace error:", resDel.Error)
		return Response{}, resDel.Error
	}

	return Response{Id: entity.Id, Name: entity.Name}, nil
}

