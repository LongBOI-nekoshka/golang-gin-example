package models

import (
	"errors"

	"angrieeecat.com/auth"
	"angrieeecat.com/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name     string
	Password string `json:"-"`
	Health   int
	Attack   int
	WeaponID int `gorm:"default:null"`
	Weapon   Weapon
}

func CheckPassword(password, hasedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))
}

func LoginCheck(name string, password string) (string, error) {

	var err error

	p := Player{}

	err = database.DB.Model(Player{}).Where("name = ?", name).Take(&p).Error

	if err != nil {
		return "", err
	}

	err = CheckPassword(password, p.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := auth.GenerateToken(p.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetPlayerByID(pid uint) (Player, error) {
	var p Player

	if err := database.DB.First(&p, pid).Error; err != nil {
		return p, errors.New("User not found")
	}

	return p, nil
}
