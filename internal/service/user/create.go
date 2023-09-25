package user

import (
	"context"
)

type (
	// CreateReq ...
	CreateReq struct {
		// add request params
	}

	// CreateResp ...
	CreateResp struct {
		// add response params
	}
)

// PostCreate ...
func (l *User) PostCreate(ctx context.Context, req *CreateReq) (*CreateResp, error) {
	// add your code here
	return &CreateResp{}, nil
}
