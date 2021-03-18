package members

import (
	"audit-poc/internal/members/models"
	"audit-poc/internal/members/payloads"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
)

type ServiceMethods interface {
	ParseMember(entity io.ReadCloser) (payloads.Request, error)
	AssociateMemberToUserGroup(ctx context.Context, request payloads.Request, groupId uuid.UUID) (payloads.Response, error)
}

type MemberRepository struct {
	db *gorm.DB
}

func (main MemberRepository) ParseMember(entity io.ReadCloser) (payloads.Request, error) {
	var newSubs *payloads.Request

	err := json.NewDecoder(entity).Decode(&newSubs)
	if err != nil {
		return payloads.Request{}, errors.New("Decode.Error")
	}

	return *newSubs, nil
}

func (main MemberRepository) AssociateMemberToUserGroup(ctx context.Context, request payloads.Request, groupId uuid.UUID) (payloads.Response, error) {
	entity := models.Member{
		Id:          uuid.New(),
		UserGroupId: groupId,
		Username:    request.Username,
	}

	err := main.AssociateWithContext(ctx, entity)
	if err != nil {
		return payloads.Response{}, err
	}

	return payloads.Response{
		Id:          entity.Id,
		Username:    entity.Username,
		UserGroupId: entity.UserGroupId,
	}, nil
}

func NewMain(db *gorm.DB) ServiceMethods {
	return MemberRepository{db}
}
