package main

import (
	"github.com/bubo-py/goonline/controllers"
	"github.com/bubo-py/goonline/database"
	"github.com/bubo-py/goonline/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MySQL database
	database.Connect("root:root@tcp(localhost:3306)/goonline?parseTime=true")
	database.Migrate()
	database.InitializeData()

	// Initialize and run router
	r := initRouter()
	r.Run()
}

// initRouter sets the routing
func initRouter() *gin.Engine {
	r := gin.Default()

	public := r.Group("/public")
	{
		public.POST("/login", controllers.GenerateToken)
		public.POST("/register", controllers.RegisterProfile)

		api := public.Group("/api").Use(middlewares.Auth())
		{
			api.POST("/logout", controllers.Logout)
			api.GET("/profile/:id", controllers.GetProfile)
			api.GET("/pokemon/:id", controllers.GetPokemon)

		}
	}
	return r
}
