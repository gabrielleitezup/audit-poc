package circle

import (
	"audit-poc/internal/circle/models"
	"audit-poc/internal/circle/payloads"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
)

type ServiceMethods interface {
	ParseCircle(entity io.ReadCloser) (payloads.Request, error)
	CreateCircle(ctx context.Context, request payloads.Request) (payloads.Response, error)
}

type CircleRepository struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) ServiceMethods {
	return CircleRepository{db}
}

func (main CircleRepository) ParseCircle(entity io.ReadCloser) (payloads.Request, error) {
	var newSubs *payloads.Request

	err := json.NewDecoder(entity).Decode(&newSubs)
	if err != nil {
		return payloads.Request{}, errors.New("Decode.Error")
	}

	return *newSubs, nil
}

func (main CircleRepository) CreateCircle(ctx context.Context, request payloads.Request) (payloads.Response, error) {
	entity := models.Circle{
		Id:          uuid.New(),
		Name:        request.Name,
		Rules:       request.Rules,
		WorkspaceId: request.WorkspaceId,
	}

	err := main.SaveWithContext(ctx, entity)
	if err != nil {
		return payloads.Response{}, err
	}

	return payloads.Response{
		Id:   entity.Id,
		Name: entity.Name,
	}, nil
}
