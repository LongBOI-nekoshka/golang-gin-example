package controllers

import (
	"log"

	"angrieeecat.com/database"
	"angrieeecat.com/database/models"
	"github.com/gin-gonic/gin"
)

type CreateWeaponInput struct {
	Name             string `json:"name" binding:"required"`
	Damage           int    `json:"damage" binding:"required"`
	HealthPercentage int    `json:"health_percentage"`
	AttackPercentage int    `json:"attack_percentage"`
	Range            int    `json:"range" binding:"required"`
	Droprate         int    `json:"droprate" binding:"required"`
	EnemyID          int    `json:"enemyID"`
}

func CreateWeapon(c *gin.Context) {

	body := &CreateWeaponInput{}

	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.Bind(&body)
	weapon := models.Weapon{
		Name:             body.Name,
		Damage:           body.Damage,
		HealthPercentage: body.HealthPercentage,
		AttackPercentage: body.AttackPercentage,
		Range:            body.Range,
		Droprate:         body.Droprate,
		EnemyID:          body.EnemyID,
	}
	result := database.DB.Create(&weapon)
	if result.Error != nil {
		c.Status(400)
		log.Fatal(result.Error)
	}
	c.JSON(200, gin.H{
		"message": weapon,
	})

}

func BulkCreation(c *gin.Context) {

}
