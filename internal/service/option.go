package service

import (
	"gin-plus-admin/internal/service/auth"
	"gin-plus-admin/internal/service/role"
	"gin-plus-admin/internal/service/user"
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

func WithAuthApi(api *auth.Auth) ApiOption {
	return func(a *Api) {
		a.Auth = api
	}
}
