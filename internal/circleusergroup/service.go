package circleusergroup

import (
	"audit-poc/internal/circleusergroup/models"
	"audit-poc/internal/circleusergroup/payloads"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
)

type ServiceMethods interface {
	ParseCircleUserGroup(entity io.ReadCloser) (payloads.Request, error)
	AssociateCircleUserGroup(ctx context.Context, request payloads.Request, circleId uuid.UUID) (payloads.Response, error)
}

type CircleUserGroupRepository struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) ServiceMethods {
	return CircleUserGroupRepository{db}
}

func (main CircleUserGroupRepository) ParseCircleUserGroup(entity io.ReadCloser) (payloads.Request, error) {
	var newSubs *payloads.Request

	err := json.NewDecoder(entity).Decode(&newSubs)
	if err != nil {
		return payloads.Request{}, errors.New("Decode.Error")
	}

	return *newSubs, nil
}

func (main CircleUserGroupRepository) AssociateCircleUserGroup(ctx context.Context, request payloads.Request, circleId uuid.UUID) (payloads.Response, error) {
	entity := models.CircleUserGroup{
		Id:          uuid.New(),
		CircleId:    circleId,
		UserGroupId: request.UserGroupId,
		DeletedAt:   gorm.DeletedAt{},
	}

	err := main.db.WithContext(ctx).Model(&models.CircleUserGroup{}).Create(&entity)
	if err.Error != nil {
		return payloads.Response{}, err.Error
	}

	return payloads.Response{
		CircleId:    entity.CircleId,
		UserGroupId: entity.UserGroupId,
	}, nil
}
