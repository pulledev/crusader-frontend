package model

import (
	"gorm.io/gorm"
)

import "time"

type Member struct {
	DiscordId string `gorm:"primaryKey;autoIncrement:false"`
	SteamId   string
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAT time.Time
}

type Element struct {
	gorm.Model
}
