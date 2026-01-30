package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/knr1997/quiz-tracker-apiserver/core"
	"github.com/knr1997/quiz-tracker-apiserver/database"
	"github.com/knr1997/quiz-tracker-apiserver/global"
	"github.com/knr1997/quiz-tracker-apiserver/model/example"
	"github.com/knr1997/quiz-tracker-apiserver/model/system"
	"github.com/knr1997/quiz-tracker-apiserver/routes"
	"go.uber.org/zap"
)

func main() {
	// Connect DB
	database.Connect()

	// Auto migrate
	err := database.DB.AutoMigrate(

		system.SysUser{},
		example.ExaCustomer{},
		// system.SysBaseMenu{},
		// system.SysAuthority{},
		// system.SysBaseMenuBtn{},
	)
	if err != nil {
		log.Fatal(err)
	}

	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)

	// Gin setup
	r := gin.Default()

	// Routes
	routes.RegisterRoutes(r)

	// Start server
	r.Run(":8080")
}
