package userworkspace

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
)

type UserGroupWorkspace struct {
	Id          uuid.UUID      `json:"id"`
	UserGroupId uuid.UUID      `json:"user_group_id"`
	WorkspaceId uuid.UUID      `json:"workspace_id"`
	Permission  []byte         `json:"permission" gorm:"type:jsonb"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

func (main Main) ParseUserWorkspace(entity io.ReadCloser) (Request, error) {
	var newSubs *Request

	err := json.NewDecoder(entity).Decode(&newSubs)
	if err != nil {
		return Request{}, errors.New("Decode.Error")
	}

	return *newSubs, nil
}

func (main Main) AssociateUserGroupToWorkspace(ctx context.Context, request Request, workspaceId uuid.UUID) (Response, error) {
	entity := UserGroupWorkspace{
		Id:          uuid.New(),
		UserGroupId: request.GroupId,
		WorkspaceId: workspaceId,
		Permission:  request.Permissions,
	}

	res := main.db.WithContext(ctx).Model(&UserGroupWorkspace{}).Create(&entity)
	if res.Error != nil {
		logrus.Errorln("Save Workspace error:", res.Error)
		return Response{}, res.Error
	}

	return Response{Id: entity.Id, GroupId: entity.UserGroupId, WorkspaceId: entity.WorkspaceId}, nil
}
