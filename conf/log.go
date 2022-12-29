package conf

import (
	"fmt"
	"os"

	"github.com/rymiyamoto/affiliate-api/log"
)

// InitLog ログ初期化
func InitLog(logTypeName string) error {
	log.SetTyp(logTypeName)
	log.SetOutput(os.Stdout)

	val, err := GetEnv("DEBUG_LOG")
	if err != nil {
		return fmt.Errorf("failed to init log. err: %w", err)
	}

	if val == "enable" {
		log.SetLevel(log.DEBUG)
	}

	return nil
}
