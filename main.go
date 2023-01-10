package main

import (
	"net/http"

	"angrieeecat.com/auth"
	"angrieeecat.com/controllers"
	"angrieeecat.com/database"
	"angrieeecat.com/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.InitEnv()
	database.ConnectToDB()
}

// middleware
func checkChar() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.IsTokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()
	player := r.Group("/player", checkChar())
	weapon := r.Group("/weapon", checkChar())
	auth := r.Group("/auth")
	chat := r.Group("/chat")

	chat.GET("/:roomId", controllers.Chat)

	player.GET("/current", controllers.GetCurrentPlayer)
	player.POST("/update", controllers.UpdatePlayer)

	weapon.POST("/create", controllers.CreateWeapon)

	auth.POST("/register", controllers.CreatePlayer)
	auth.POST("/login", controllers.Login)
	r.GET("/AHHH", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": "AHHHH"})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.Run()
}
