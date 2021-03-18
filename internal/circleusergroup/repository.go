package circle

import (
	"audit-poc/internal/auditions"
	"audit-poc/internal/circleusergroup/models"
	"audit-poc/util"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (main CircleUserGroupRepository) SaveWithContext(ctx context.Context, entity models.CircleUserGroup) error {
	return main.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&models.CircleUserGroup{}).Create(&entity).Error; err != nil {
			logrus.Errorln("Associate circle to User group error:", err)
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
			logrus.Errorln("Associate circle to user group audit error:", err)
			return err
		}

		return nil
	})
}
