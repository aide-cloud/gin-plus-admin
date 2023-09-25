package role

import (
	"context"
)

type (
	// ListReq ...
	ListReq struct {
		// add request params
	}

	// ListResp ...
	ListResp struct {
		// add response params
	}
)

// List ...
func (l *Role) List(ctx context.Context, req *ListReq) (*ListResp, error) {
	// add your code here
	return &ListResp{}, nil
}
