package auth

import (
	"context"
)

type (
	// LogoutService ...
	LogoutParams struct {
		// add params
	}

	// LogoutReply ...
	LogoutReply struct {
		// add reply
	}
)

// Logout ...
func (l *Auth) Logout(ctx context.Context, req *LogoutParams) (*LogoutReply, error) {
	// add your code here
	return &LogoutReply{}, nil
}
