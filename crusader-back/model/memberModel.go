package model

import (
	"gorm.io/gorm"
)

import "time"

type Unit struct {
	gorm.Model
	Name          string
	Description   string
	DiscordRoleId string
	UnitRoles     []UnitRoles
}

type UnitRoles struct {
	gorm.Model
	Name        string `gorm:"unique;"`
	Description string
}

type Member struct {
	DiscordId      string `gorm:"primaryKey;autoIncrement:false"`
	Name           string
	SteamId        string
	Unit           Unit           //one-to-Many
	MembershipType MembershipType //one-to-Many
	Rank           Rank           //one-to-many
	Stab           Stab           //many-to-many
	DiscordNick    string
	CreatedAt      time.Time
	UpdatedAT      time.Time
}

type Element struct {
	gorm.Model
}
