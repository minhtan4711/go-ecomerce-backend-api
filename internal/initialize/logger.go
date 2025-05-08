package initialize

import (
	"github.com/go-ecomerce-backend-api/global"
	"github.com/go-ecomerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
