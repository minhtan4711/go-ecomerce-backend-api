package initialize

import (
	"fmt"

	"github.com/go-ecomerce-backend-api/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration mysql", global.Config.Mysql.Host)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run()
}
