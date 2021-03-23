package auditions

import (
	"audit-poc/internal/auditions/models"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type ServiceMethods interface {
	HistoryList(ctx context.Context, params map[string]interface{}) ([]models.Audition, error)
}

type AuditionRepository struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) ServiceMethods {
	return AuditionRepository{db}
}

func (main AuditionRepository) HistoryList(ctx context.Context, params map[string]interface{}) ([]models.Audition, error) {
	var entity []models.Audition

	res := main.db.WithContext(ctx).Model(&models.Audition{}).Where(params).Find(&entity)
	if res.Error != nil {
		return []models.Audition{}, res.Error
	}

	for key, value := range params {
		fmt.Println(key, " : ", value)
	}

	return entity, nil
}


