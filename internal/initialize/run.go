package initialize

import (
	"fmt"

	"github.com/go-ecomerce-backend-api/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration mysql", global.Config.Mysql.Host)
	InitLogger()
	global.Logger.Logger.Info("Config log ok")
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run()
}
