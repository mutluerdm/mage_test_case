package main

import (
	"context"
	"fmt"
	"mage_test_case/config"
	"mage_test_case/controller"
	"mage_test_case/mlog"
)

func main() {
	fmt.Println("hola\nAPI UP!")
	mlog.InitLoggers()
	conf := config.LoadConfigs()
	mlog.Printf("config loaded : %+v", conf.Title)
	apiController, err := controller.NewAPI(&conf)
	if err != nil {
		mlog.PrintErrf("Server Cannot Start err : %+v", err)
		return
	}
	apiController.Start(conf, context.Background())
}
