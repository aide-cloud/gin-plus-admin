package role

import (
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ IRole = (*Role)(nil)

type (
	// IRole ...
	IRole interface {
		ginplus.Middlewarer
		ginplus.MethoderMiddlewarer
		// add interface method
	}

	// Role ...
	Role struct {
		// add child module
	}

	// RoleOption ...
	RoleOption func(*Role)
)

// defaultRole ...
func defaultRole() *Role {
	return &Role{
		// add child module
	}
}

// NewRole ...
func NewRole(opts ...RoleOption) *Role {
	m := defaultRole()
	for _, o := range opts {
		o(m)
	}

	return m
}

// Middlewares ...
func (l *Role) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			// your code ...
		},
		// other ...
	}
}

// MethoderMiddlewares ...
func (l *Role) MethoderMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		// your method middlewares
	}
}
