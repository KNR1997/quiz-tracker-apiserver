package global

import (
	"github.com/knr1997/quiz-tracker-apiserver/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB
	GVA_LOG    *zap.Logger
	GVA_CONFIG config.Server
)
