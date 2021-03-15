package usergroup

import (
	"audit-poc/internal/auditions"
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (u *UserGroup) AfterCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context

	cs, err := json.Marshal(u)
	if err != nil {
		return err
	}

	audit := auditions.Audition{
		Id:           uuid.New(),
		Username:     ctx.Value("jwt").(string),
		TableName:    tx.Statement.Table,
		Operation:    "INSERT",
		EntityId:     u.Id.String(),
		CurrentState: cs,
		UserIpAddr:   ctx.Value("user-ip").(string),
		UserAgent:    ctx.Value("user-agent").(string),
	}

	svAudit := tx.Model(&auditions.Audition{}).Create(&audit)
	if svAudit.Error != nil {
		return svAudit.Error
	}

	return nil
}

func (u *UserGroup) AfterUpdate(tx *gorm.DB) error {
	ctx := tx.Statement.Context

	cs, err := json.Marshal(u)
	if err != nil {
		return err
	}

	audit := auditions.Audition{
		Id:           uuid.New(),
		Username:     ctx.Value("jwt").(string),
		TableName:    tx.Statement.Table,
		Operation:    "UPDATE",
		EntityId:     u.Id.String(),
		CurrentState: cs,
		UserIpAddr:   ctx.Value("user-ip").(string),
		UserAgent:    ctx.Value("user-agent").(string),
	}

	svAudit := tx.Model(&auditions.Audition{}).Create(&audit)
	if svAudit.Error != nil {
		return svAudit.Error
	}

	return nil
}

func (u *UserGroup) AfterDelete(tx *gorm.DB) error {
	ctx := tx.Statement.Context

	cs, err := json.Marshal(u)
	if err != nil {
		return err
	}


	audit := auditions.Audition{
		Id:           uuid.New(),
		Username:     ctx.Value("jwt").(string),
		TableName:    tx.Statement.Table,
		Operation:    "DELETE",
		EntityId:     u.Id.String(),
		CurrentState: cs,
		UserIpAddr:   ctx.Value("user-ip").(string),
		UserAgent:    ctx.Value("user-agent").(string),
	}

	svAudit := tx.Model(&auditions.Audition{}).Create(&audit)
	if svAudit.Error != nil {
		return svAudit.Error
	}

	return nil
}