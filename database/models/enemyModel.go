package models

import "gorm.io/gorm"

type Enemy struct {
	gorm.Model
	Name       string
	Health     int
	Attack     int
	IsBoss     bool
	IsMiniBoss bool
}
