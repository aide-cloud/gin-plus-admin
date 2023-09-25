package role

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

// Create ...
func (l *Role) Create(ctx context.Context, req *CreateReq) (*CreateResp, error) {
	// add your code here
	return &CreateResp{}, nil
}
