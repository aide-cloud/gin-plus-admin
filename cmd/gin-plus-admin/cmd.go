package main

import (
	"flag"
	"gin-plus-admin/internal/server"

	ginplus "github.com/aide-cloud/gin-plus"
)

var (
	// Version 版本号
	Version string
	// ServiceName 服务名称
	ServiceName = "gin-plus-admin"
	// configPath 配置文件路径
	configPath = flag.String("config", "config/config.yaml", "config file path")
)

func main() {
	flag.Parse()
	bc := Init()

	svs := []ginplus.Server{
		server.NewHttpServer(bc.Server),
	}

	// 启动gin-plus
	ginplus.NewCtrlC(svs...).Start()
}
