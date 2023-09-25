package user

import (
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ IUser = (*User)(nil)

type (
	// IUser ...
	IUser interface {
		ginplus.Middlewarer
		ginplus.MethoderMiddlewarer
		// add interface method
	}

	// User ...
	User struct {
		// add child module
	}

	// UserOption ...
	UserOption func(*User)
)

// defaultUser ...
func defaultUser() *User {
	return &User{
		// add child module
	}
}

// NewUser ...
func NewUser(opts ...UserOption) *User {
	m := defaultUser()
	for _, o := range opts {
		o(m)
	}

	return m
}

// Middlewares ...
func (l *User) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			// your code ...
		},
		// other ...
	}
}

// MethoderMiddlewares ...
func (l *User) MethoderMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		// your method middlewares
	}
}
