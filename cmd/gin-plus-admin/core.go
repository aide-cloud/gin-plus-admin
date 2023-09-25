package main

import (
	"gin-plus-admin/internal/conf"
	"gin-plus-admin/pkg/conn"

	ginplus "github.com/aide-cloud/gin-plus"
)

func Init() {
	bc := conf.NewBootstrap(conf.WithConfigFile(*configPath))
	if bc == nil {
		panic("conf.NewBootstrap() error")
	}

	if bc.Server != nil && bc.Server.Name != "" {
		ServiceName = bc.Server.Name
	}

	ginplus.Logger().Sugar().Infof("%s version: %s", ServiceName, Version)

	if bc.Data != nil {
		if bc.Data.Mysql != nil {
			conn.InitMysqlDB(bc.Data.Mysql.DSN, bc.Data.Mysql.Debug)
		}
	}
}

var httpMethodPrefixes = []ginplus.HttpMethod{
	{
		Prefix: "Create",
		Method: ginplus.Post,
	}, {
		Prefix: "Update",
		Method: ginplus.Put,
	}, {
		Prefix: "Edit",
		Method: ginplus.Put,
	}, {
		Prefix: "Delete",
		Method: ginplus.Delete,
	}, {
		Prefix: "Detail",
		Method: ginplus.Get,
	}, {
		Prefix: "List",
		Method: ginplus.Get,
	},
}
