package conf

import (
	"fmt"
	"os"

	"github.com/rymiyamoto/affiliate-api/env"
	"github.com/rymiyamoto/affiliate-api/log"
	"github.com/rymiyamoto/affiliate-api/model/common"
)

// InitLog ログ初期化
func InitLog(logTypeName string) error {
	log.SetTyp(logTypeName)
	log.SetOutput(os.Stdout)

	val, err := env.Value("GO_ENV")
	if err != nil {
		return fmt.Errorf("failed to init log. err: %w", err)
	}

	if val == common.EnvDev {
		log.SetLevel(log.DEBUG)
	}

	return nil
}
