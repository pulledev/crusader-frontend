package model

import (
	"time"

	"gorm.io/gorm"
)

type Unit struct {
	gorm.Model
	Name          string
	Description   string
	DiscordRoleId string
	UnitRole      []UnitRole `gorm:"many2many:unit_unitRoles;"`
}

type UnitRole struct {
	gorm.Model
	Name        string `gorm:"unique;"`
	Description string
}

type MembershipType struct {
	gorm.Model
	Name string
}

type Rank struct {
	Level     int `gorm:"primaryKey;autoIncrement:false"`
	Name      string
	CreatedAt time.Time
	UpdatedAT time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Stab struct {
	gorm.Model
	Name        string
	Description string
}

type Member struct {
	DiscordId        string `gorm:"primaryKey;autoIncrement:false"`
	Name             string
	SteamId          string
	UnitID           int            `gorm:"default:0"`
	Unit             Unit           `gorm:"foreignKey:UnitID"`
	MembershipTypeID int            `gorm:"default:0"`
	MembershipType   MembershipType `gorm:"foreignKey:MembershipTypeID"`
	RankLevel        int            `gorm:"default:0"`
	Rank             Rank           `gorm:"foreignKey:RankLevel;default:0"`
	Stab             []Stab         `gorm:"many2many:member_stab;"`
	DiscordNick      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
