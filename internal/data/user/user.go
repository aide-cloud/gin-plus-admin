package user

import (
	"gin-plus-admin/pkg/conn"
	"gin-plus-admin/pkg/model"
	"gin-plus-admin/pkg/model/query"
)

type (
	// User ...
	User struct {
		query.IAction[model.User]
	}
)

// NewUser ...
func NewUser() *User {
	return &User{
		IAction: query.NewAction(query.WithDB[model.User](conn.GetMysqlDB())),
	}
}
