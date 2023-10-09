package auth

import (
	"context"
)

type (
	// LoginService ...
	LoginParams struct {
		// add params
	}

	// LoginReply ...
	LoginReply struct {
		// add reply
	}
)

// Login ...
func (l *Auth) Login(ctx context.Context, req *LoginParams) (*LoginReply, error) {
	// add your code here
	return &LoginReply{}, nil
}
