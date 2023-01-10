package models

import "gorm.io/gorm"

type BoardCell struct {
	gorm.Model
	SnakePercent  int
	LadderPercent int
	HaveBoss      bool
}
