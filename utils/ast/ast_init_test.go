package ast

import (
	"path/filepath"

	"github.com/knr1997/quiz-tracker-apiserver/global"
)

func init() {
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("../../../")
	global.GVA_CONFIG.AutoCode.Server = "server"
}
