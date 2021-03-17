package usergroup

import (
	"audit-poc/internal/auditions"
	"audit-poc/internal/userworkspace/usergroup/models"
	"audit-poc/internal/userworkspace/usergroup/payloads"
	"audit-poc/util"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (main UserGroupRepository) SaveWithContext(ctx context.Context, entity models.UserGroup) error {
	return main.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&models.UserGroup{}).Create(&entity).Error; err != nil {
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

func (main UserGroupRepository) UpdateWithContext(ctx context.Context, entity models.UserGroup, request payloads.Request, id uuid.UUID) (models.UserGroup, error) {
	return entity, main.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&models.UserGroup{}).Where("id", id).Update("name", request.Name).Scan(&entity).Error; err != nil {
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

func (main UserGroupRepository) DeleteWithContext(ctx context.Context, entity models.UserGroup, id uuid.UUID)  error {
	return main.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&models.UserGroup{}).First(&entity, id).Delete(&entity, id).Error; err != nil {
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