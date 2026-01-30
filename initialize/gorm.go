package initialize

import (
	"os"

	"github.com/knr1997/quiz-tracker-apiserver/model/system"
	"gorm.io/gorm"
)

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(

		system.SysUser{},
		system.SysBaseMenu{},
		system.SysAuthority{},
		system.SysBaseMenuBtn{},
	)
	if err != nil {
		os.Exit(0)
	}

}
