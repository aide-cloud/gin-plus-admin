package user

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

// PostList ...
func (l *User) List(ctx context.Context, req *ListReq) (*ListResp, error) {
	// add your code here
	return &ListResp{}, nil
}
