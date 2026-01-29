package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/knr1997/quiz-tracker-apiserver/database"
	"github.com/knr1997/quiz-tracker-apiserver/models"
	"github.com/knr1997/quiz-tracker-apiserver/routes"
)

func main() {
	// Connect DB
	database.Connect()

	// Auto migrate
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	// Gin setup
	r := gin.Default()

	// Routes
	routes.RegisterRoutes(r)

	// Start server
	r.Run(":8080")
}
