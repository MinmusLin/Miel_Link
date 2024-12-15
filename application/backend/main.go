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
		fmt.Printf("Init settings failed, error: %v\n", err)
	}
	if err := pkg.MysqlInit(); err != nil {
		fmt.Printf("Init mysql failed, error: %v\n", err)
	}
	r := router.SetupRouter()
	_ = r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))
}
