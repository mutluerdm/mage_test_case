package main

import (
	"fmt"
	"mage_test_case/config"
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
	//repo := repo.New(conf)
	//serverIns := model.ServerInstanceConfig{Ip: conf.Nebula.Host, Port: conf.Nebula.Port, Active: true}
	//gameController := controller.NewGame(repo, &serverIns, geoipDb)
	//gameController.Start()

}
