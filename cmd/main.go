package main

import (
	"flag"

	"gitee.com/haiyux/kratos-layout/internal/config"
	"gitee.com/haiyux/kratos-layout/internal/logging"
	"gitee.com/haiyux/kratos-layout/internal/server"
)

var conf string

func init() {
	flag.StringVar(&conf, "conf", "config/", "config path, eg: -conf config/")
}

func main() {
	flag.Parse()
	var a = new(App)
	if err := config.InitConfig(conf, a); err != nil {
		panic(err)
	}
	log := logging.InitLogger(a.LogPath)
	app, err := server.InitServer(a.Addr, log)
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}
}

type App struct {
	Addr    string `json:"addr"`
	LogPath string `json:"log_path"`
}
