package models

import (
	"audit-poc/internal/auditions"
	"audit-poc/util"
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (d *Deployment) AfterCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context

	cs, err := json.Marshal(d)
	if err != nil {
		return err
	}

	audit := auditions.Audition{
		Id:           uuid.New(),
		Username:     ctx.Value(util.AuthContextKey).(string),
		TableName:    tx.Statement.Table,
		Operation:    "INSERT",
		EntityId:     d.Id.String(),
		CurrentState: cs,
		UserIpAddr:   ctx.Value(util.UserIpContextKey).(string),
		UserAgent:    ctx.Value(util.UserAgentContextKey).(string),
	}

	svAudit := tx.Model(&auditions.Audition{}).Create(&audit)
	if svAudit.Error != nil {
		return svAudit.Error
	}

	return nil
}

func (d *Deployment) AfterUpdate(tx *gorm.DB) error {
	ctx := tx.Statement.Context

	cs, err := json.Marshal(d)
	if err != nil {
		return err
	}

	audit := auditions.Audition{
		Id:           uuid.New(),
		Username:     ctx.Value(util.AuthContextKey).(string),
		TableName:    tx.Statement.Table,
		Operation:    "UPDATE",
		EntityId:     d.Id.String(),
		CurrentState: cs,
		UserIpAddr:   ctx.Value(util.UserIpContextKey).(string),
		UserAgent:    ctx.Value(util.UserAgentContextKey).(string),
	}

	svAudit := tx.Model(&auditions.Audition{}).Create(&audit)
	if svAudit.Error != nil {
		return svAudit.Error
	}

	return nil
}

func (d *Deployment) AfterDelete(tx *gorm.DB) error {
	ctx := tx.Statement.Context

	cs, err := json.Marshal(d)
	if err != nil {
		return err
	}

	audit := auditions.Audition{
		Id:           uuid.New(),
		Username:     ctx.Value(util.AuthContextKey).(string),
		TableName:    tx.Statement.Table,
		Operation:    "DELETE",
		EntityId:     d.Id.String(),
		CurrentState: cs,
		UserIpAddr:   ctx.Value(util.UserIpContextKey).(string),
		UserAgent:    ctx.Value(util.UserAgentContextKey).(string),
	}

	svAudit := tx.Model(&auditions.Audition{}).Create(&audit)
	if svAudit.Error != nil {
		return svAudit.Error
	}

	return nil
}
