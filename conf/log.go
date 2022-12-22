package conf

import (
	"io"
	"os"

	"github.com/rymiyamoto/affiliate-api/log"
)

// APILog apiログの出力先
func APILog() io.Writer {
	return os.Stdout
}

// InitLog は、ログ初期化を行う
func InitLog(logTypeName string) {
	log.SetTyp(logTypeName)
	log.SetOutput(APILog())
	if DebugEnabled() {
		log.SetLevel(log.DEBUG)
	}
}

// DebugEnabled DEBUGレベルのログ出力が有効かどうか
func DebugEnabled() bool {
	if v := os.Getenv("DEBUG_LOG"); v == "enable" {
		return true
	}
	return false
}
