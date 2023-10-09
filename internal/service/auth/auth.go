package auth

import (
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ IAuth = (*Auth)(nil)

type (
	// IAuth ...
	IAuth interface {
		ginplus.Middlewarer
		ginplus.MethoderMiddlewarer
		// add interface method
	}

	// Auth ...
	Auth struct {
		// add child module
	}

	// AuthOption ...
	AuthOption func(*Auth)
)

// defaultAuth ...
func defaultAuth() *Auth {
	return &Auth{
		// add child module
	}
}

// NewAuth ...
func NewAuth(opts ...AuthOption) *Auth {
	m := defaultAuth()
	for _, o := range opts {
		o(m)
	}

	return m
}

// Middlewares ...
func (l *Auth) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			// your code ...
		},
		// other ...
	}
}

// MethoderMiddlewares ...
func (l *Auth) MethoderMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		// your method middlewares
	}
}
