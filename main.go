package main

import (
	"context"
	"fmt"
	"mage_test_case/config"
	"mage_test_case/controller"
	"mage_test_case/mlog"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("hola\n API UP!")
	mlog.InitLoggers()
	conf := config.LoadConfigs()
	mlog.Printf("config loaded : %+v", conf.Title)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		mlog.Fatal(err)
	}
	mlog.Println("root dir : ", dir)
	ac := controller.NewAPI(&conf)
	ac.Start(conf, context.Background())

}
