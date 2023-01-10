package models

import "gorm.io/gorm"

type Weapon struct {
	gorm.Model
	Name             string
	Damage           int `gorm:"default:null"`
	HealthPercentage int
	AttackPercentage int
	Range            int
	Droprate         int
	EnemyID          int `gorm:"default:null"`
	Enemy            Enemy
}
