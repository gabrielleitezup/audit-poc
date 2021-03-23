package models

import (
	"audit-poc/internal/auditions"
	"audit-poc/util"
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (uw *UserGroupWorkspace) AfterCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context

	cs, err := json.Marshal(uw)
	if err != nil {
		return err
	}

	audit := auditions.Audition{
		Id:           uuid.New(),
		Username:     ctx.Value(util.AuthContextKey).(string),
		TableName:    tx.Statement.Table,
		Operation:    "INSERT",
		EntityId:     uw.Id.String(),
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

func (uw *UserGroupWorkspace) AfterUpdate(tx *gorm.DB) error {
	ctx := tx.Statement.Context

	cs, err := json.Marshal(uw)
	if err != nil {
		return err
	}

	audit := auditions.Audition{
		Id:           uuid.New(),
		Username:     ctx.Value(util.AuthContextKey).(string),
		TableName:    tx.Statement.Table,
		Operation:    "UPDATE",
		EntityId:     uw.Id.String(),
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

func (uw *UserGroupWorkspace) AfterDelete(tx *gorm.DB) error {
	ctx := tx.Statement.Context

	cs, err := json.Marshal(uw)
	if err != nil {
		return err
	}


	audit := auditions.Audition{
		Id:           uuid.New(),
		Username:     ctx.Value(util.AuthContextKey).(string),
		TableName:    tx.Statement.Table,
		Operation:    "DELETE",
		EntityId:     uw.Id.String(),
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