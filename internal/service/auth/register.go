package auth

import (
	"context"
)

type (
	// RegisterService ...
	RegisterParams struct {
		// add params
	}

	// RegisterReply ...
	RegisterReply struct {
		// add reply
	}
)

// Register ...
func (l *Auth) Register(ctx context.Context, req *RegisterParams) (*RegisterReply, error) {
	// add your code here
	return &RegisterReply{}, nil
}
