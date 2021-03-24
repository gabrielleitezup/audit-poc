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

	startedAt, okS := params["startedAt"]
	endedAt, okE := params["endedAt"]

	if okS == true && okE == true {
		delete(params, "startedAt")
		delete(params, "endedAt")
		res := main.db.WithContext(ctx).Model(&models.Audition{}).Where(params).Where("created_at BETWEEN ? AND ?", startedAt, endedAt).Find(&entity)
		if res.Error != nil {
			return []models.Audition{}, res.Error
		}
	} else {
		res := main.db.WithContext(ctx).Model(&models.Audition{}).Where(params).Find(&entity)
		if res.Error != nil {
			return []models.Audition{}, res.Error
		}
	}

	for key, value := range params {
		fmt.Println(key, " : ", value)
	}

	return entity, nil
}
