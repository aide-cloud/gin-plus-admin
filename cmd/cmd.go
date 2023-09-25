package main

import (
	"gin-plus-admin/internal/api"
	"gin-plus-admin/internal/api/logic/role"

	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化gin实例
	r := gin.Default()
	r.Use()

	// 初始化ginplus实例
	ginplusEngine := ginplus.New(r,
		ginplus.AppendHttpMethodPrefixes(httpMethodPrefixes...),
		// 注册api模块
		ginplus.WithControllers(
			api.NewApi(
				// api.WithUserApi(user.NewUser()),
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
