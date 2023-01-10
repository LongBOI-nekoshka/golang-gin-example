package main

import (
	"angrieeecat.com/database"
	"angrieeecat.com/database/models"
	"angrieeecat.com/initializers"
)

func init() {
	initializers.InitEnv()
	database.ConnectToDB()
}

func main() {
	database.DB.AutoMigrate(models.Player{}, models.Enemy{}, models.Weapon{}, models.BoardCell{})
}
