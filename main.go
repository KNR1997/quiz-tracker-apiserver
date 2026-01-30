package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/knr1997/quiz-tracker-apiserver/database"
	"github.com/knr1997/quiz-tracker-apiserver/model/system"
	"github.com/knr1997/quiz-tracker-apiserver/routes"
)

func main() {
	// Connect DB
	database.Connect()

	// Auto migrate
	err := database.DB.AutoMigrate(

		system.SysUser{},
		// system.SysBaseMenu{},
		// system.SysAuthority{},
		// system.SysBaseMenuBtn{},
	)
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
