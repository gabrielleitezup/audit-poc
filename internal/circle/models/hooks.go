package models

import (
	"audit-poc/internal/auditions/models"
	"audit-poc/util"
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c *Circle) AfterCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context

	cs, err := json.Marshal(c)
	if err != nil {
		return err
	}

	audit := models.Audition{
		Id:           uuid.New(),
		Username:     ctx.Value(util.AuthContextKey).(string),
		TableName:    tx.Statement.Table,
		Operation:    "INSERT",
		EntityId:     c.Id.String(),
		CurrentState: cs,
		UserIpAddr:   ctx.Value(util.UserIpContextKey).(string),
		UserAgent:    ctx.Value(util.UserAgentContextKey).(string),
	}

	svAudit := tx.Model(&models.Audition{}).Create(&audit)
	if svAudit.Error != nil {
		return svAudit.Error
	}

	return nil
}

func (c *Circle) AfterUpdate(tx *gorm.DB) error {
	ctx := tx.Statement.Context

	cs, err := json.Marshal(c)
	if err != nil {
		return err
	}

	audit := models.Audition{
		Id:           uuid.New(),
		Username:     ctx.Value(util.AuthContextKey).(string),
		TableName:    tx.Statement.Table,
		Operation:    "UPDATE",
		EntityId:     c.Id.String(),
		CurrentState: cs,
		UserIpAddr:   ctx.Value(util.UserIpContextKey).(string),
		UserAgent:    ctx.Value(util.UserAgentContextKey).(string),
	}

	svAudit := tx.Model(&models.Audition{}).Create(&audit)
	if svAudit.Error != nil {
		return svAudit.Error
	}

	return nil
}

func (c *Circle) AfterDelete(tx *gorm.DB) error {
	ctx := tx.Statement.Context

	cs, err := json.Marshal(c)
	if err != nil {
		return err
	}

	audit := models.Audition{
		Id:           uuid.New(),
		Username:     ctx.Value(util.AuthContextKey).(string),
		TableName:    tx.Statement.Table,
		Operation:    "DELETE",
		EntityId:     c.Id.String(),
		CurrentState: cs,
		UserIpAddr:   ctx.Value(util.UserIpContextKey).(string),
		UserAgent:    ctx.Value(util.UserAgentContextKey).(string),
	}

	svAudit := tx.Model(&models.Audition{}).Create(&audit)
	if svAudit.Error != nil {
		return svAudit.Error
	}

	return nil
}
