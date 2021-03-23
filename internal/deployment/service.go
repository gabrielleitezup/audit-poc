package deployment

import (
	"audit-poc/internal/deployment/models"
	"audit-poc/internal/deployment/payloads"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
)

type ServiceMethods interface {
	ParseDeployment(entity io.ReadCloser) (payloads.Request, error)
	CreateDeployment(ctx context.Context, request payloads.Request, circleId uuid.UUID) (payloads.Response, error)
}

type DeploymentRepository struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) ServiceMethods {
	return DeploymentRepository{db}
}

func (main DeploymentRepository) ParseDeployment(entity io.ReadCloser) (payloads.Request, error) {
	var newSubs *payloads.Request

	err := json.NewDecoder(entity).Decode(&newSubs)
	if err != nil {
		return payloads.Request{}, errors.New("Decode.Error")
	}

	return *newSubs, nil
}

func (main DeploymentRepository) CreateDeployment(ctx context.Context, request payloads.Request, circleId uuid.UUID) (payloads.Response, error) {
	entity := models.Deployment{
		Id:       uuid.New(),
		Name:     request.Name,
		Version:  request.Version,
		CircleId: circleId,
	}

	err := main.db.WithContext(ctx).Model(&models.Deployment{}).Create(&entity)
	if err.Error != nil {
		return payloads.Response{}, err.Error
	}

	return payloads.Response{
		Id:      entity.Id,
		Name:    entity.Name,
		Version: entity.Version,
	}, nil
}
