package main

import (
	"gin-plus-admin/internal/api"
	"gin-plus-admin/internal/api/logic/role"
	"gin-plus-admin/internal/api/logic/user"
	"time"

	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var (
	// Version 版本号
	Version = "1.0.0"
	// ServiceName 服务名称
	ServiceName = "gin-plus-admin"
)

func main() {

	middle := ginplus.NewMiddleware()

	// 初始化gin实例
	r := gin.Default()
	r.Use(
		middle.Cors(),
		middle.Logger(ServiceName, time.DateTime),
	)

	// 初始化ginplus实例
	ginplusEngine := ginplus.New(r,
		ginplus.AppendHttpMethodPrefixes(httpMethodPrefixes...),
		// 注册api模块
		ginplus.WithControllers(
			api.NewApi(
				api.WithUserApi(user.NewUser()),
				api.WithRoleApi(role.NewRole()),
			),
		),
	)
	// 注册默认路由
	ginplusEngine.RegisterPing().RegisterSwaggerUI()
	// 启动gin-plus
	ginplus.NewCtrlC(ginplusEngine).Start()
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
