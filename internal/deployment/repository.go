package deployment

import (
	"audit-poc/internal/auditions"
	"audit-poc/internal/deployment/models"
	"audit-poc/util"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (main DeploymentRepository) SaveWithContext(ctx context.Context, entity models.Deployment) error {
	return main.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&models.Deployment{}).Create(&entity).Error; err != nil {
			logrus.Errorln("Create Deploy error:", err)
			return err
		}

		cs, err := json.Marshal(entity)
		if err != nil {
			return err
		}

		audit := auditions.Audition{
			Id:           uuid.New(),
			Username:     ctx.Value(util.AuthContextKey).(string),
			TableName:    ctx.Value(util.EntityContextKey).(string),
			Operation:    "INSERT",
			EntityId:     entity.Id.String(),
			CurrentState: cs,
			UserIpAddr:   ctx.Value(util.UserIpContextKey).(string),
			UserAgent:    ctx.Value(util.UserAgentContextKey).(string),
		}

		if err := tx.Model(&auditions.Audition{}).Create(&audit).Error; err != nil {
			logrus.Errorln("Create Deploy audit error:", err)
			return err
		}

		return nil
	})
}
