package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
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
		system.SysCourse{},
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

	// ✅ CORS CONFIG
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:5173",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Routes
	routes.RegisterRoutes(r)

	// Start server
	r.Run(":8080")
}
