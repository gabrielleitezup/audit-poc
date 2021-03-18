package workspace

import (
	"audit-poc/internal/auditions"
	"audit-poc/internal/workspace/models"
	"audit-poc/internal/workspace/payloads"
	"audit-poc/util"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (main WorkspaceRepository) SaveWithContext(ctx context.Context, entity models.Workspace) error {
	return main.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&models.Workspace{}).Create(&entity).Error; err != nil {
			logrus.Errorln("Save Workspace error:", err)
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
			logrus.Errorln("Save Workspace audit error:", err)
			return err
		}

		return nil
	})
}

func (main WorkspaceRepository) UpdateWithContext(ctx context.Context, entity models.Workspace, request payloads.Request, id uuid.UUID) (models.Workspace, error) {
	return entity, main.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&models.Workspace{}).Where("id", id).Update("name", request.Name).Scan(&entity).Error; err != nil {
			logrus.Errorln("Update Workspace error:", err)
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
			Operation:    "UPDATE",
			EntityId:     entity.Id.String(),
			CurrentState: cs,
			UserIpAddr:   ctx.Value(util.UserIpContextKey).(string),
			UserAgent:    ctx.Value(util.UserAgentContextKey).(string),
		}

		if err := tx.Model(&auditions.Audition{}).Create(&audit).Error; err != nil {
			logrus.Errorln("Save Workspace audit error:", err)
			return err
		}

		return nil
	})
}

func (main WorkspaceRepository) DeleteWithContext(ctx context.Context, entity models.Workspace, id uuid.UUID)  error {
	return main.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&models.Workspace{}).First(&entity, id).Delete(&entity).Error; err != nil {
			logrus.Errorln("Delete Workspace error:", err)
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
			Operation:    "DELETE",
			EntityId:     entity.Id.String(),
			CurrentState: cs,
			UserIpAddr:   ctx.Value(util.UserIpContextKey).(string),
			UserAgent:    ctx.Value(util.UserAgentContextKey).(string),
		}

		if err := tx.Model(&auditions.Audition{}).Create(&audit).Error; err != nil {
			logrus.Errorln("Save Workspace audit error:", err)
			return err
		}

		return nil
	})
}
