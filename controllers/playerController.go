package controllers

import (
	"net/http"

	"angrieeecat.com/auth"
	"angrieeecat.com/database"
	"angrieeecat.com/database/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreatePlayerInput struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdatePlayerStatsInput struct {
	Health int `json:"health" binding:"required"`
	Attack int `json:"attack" binding:"required"`
}

type UpdatePlayerPasswordInput struct {
	ID       int    `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetCurrentPlayer(c *gin.Context) {
	user_id, err := auth.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := models.GetPlayerByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": p})
}

func CreatePlayer(c *gin.Context) {
	body := &CreatePlayerInput{}

	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.Bind(&body)
	player := models.Player{
		Name:     body.Name,
		Password: EncrypPassword(body.Password),
		Health:   100,
		Attack:   10,
		WeaponID: 0,
	}
	result := database.DB.Create(&player)
	if result.Error != nil {
		c.Status(400)
	}
	c.JSON(200, gin.H{
		"message": player,
	})
}

func UpdatePlayer(c *gin.Context) {

	var player models.Player
	body := &UpdatePlayerStatsInput{}
	user_id, _ := auth.ExtractTokenID(c)

	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.Bind(&body)

	database.DB.First(&player, user_id)
	player.Health = body.Health
	player.Attack = body.Attack
	database.DB.Save(&player)
	c.JSON(200, gin.H{"message": "success", "data": player})

}

func UpdatePlayerPassword(c *gin.Context) {
	var player models.Player
	body := *&UpdatePlayerPasswordInput{}
	user_id, _ := auth.ExtractTokenID(c)

	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(400, err.Error())
		return
	}

	database.DB.First(&player, user_id)
	player.Password = EncrypPassword(body.Password)
	database.DB.Save(&player)

	c.Bind(&body)
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := models.Player{}

	p.Name = input.Name
	p.Password = input.Password

	token, err := models.LoginCheck(p.Name, p.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func EncrypPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "Something went wrong"
	}
	return string(hashedPassword)
}
