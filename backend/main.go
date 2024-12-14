package main

import (
	"backend/pkg"
	"backend/router"
	setting "backend/settings"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	if err := setting.Init(); err != nil {
		fmt.Printf("init settings failed,err:%v\n", err)
	}
	if err := pkg.MysqlInit(); err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
	}
	r := router.SetupRouter()
	_ = r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))
}
