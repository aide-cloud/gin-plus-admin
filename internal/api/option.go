package api

import (
	"gin-plus-admin/internal/api/logic/role"
	"gin-plus-admin/internal/api/logic/user"
)

func WithUserApi(api *user.User) ApiOption {
	return func(a *Api) {
		a.User = api
	}
}

func WithRoleApi(api *role.Role) ApiOption {
	return func(a *Api) {
		a.Role = api
	}
}
