package model

import (
	"gorm.io/gorm"
)

import "time"

type Member struct {
	discordId string `gorm:"primaryKey;autoIncrement:false"`
	steamId   string
	name      string
	Element   Element
	CreatedAt time.Time
	UpdatedAT time.Time
}

type Element struct {
	gorm.Model
}
