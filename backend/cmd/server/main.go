package main

import (
	"goTask/internal/controllers"
	"goTask/internal/generationService"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
	}))

	service := generationService.NewPasswordGenerationService()
	passwordGeneratorController := controllers.NewPasswordGenerationController(service)

	api := router.Group("/api/password")
	{
		api.POST("/generate", passwordGeneratorController.Generate)
		api.GET("/generationOptions", passwordGeneratorController.GetGenerationOptions)
	}

	router.Run(":8080")
}
