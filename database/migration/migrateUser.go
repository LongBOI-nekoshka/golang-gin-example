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

func migratePlayer() {
	database.DB.AutoMigrate(models.Player{})
}
