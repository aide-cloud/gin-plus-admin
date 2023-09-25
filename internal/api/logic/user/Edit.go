package user

import (
	"context"
)

type (
	// EditReq ...
	EditReq struct {
		// add request params
		ID int `uri:"id"`
	}

	// EditResp ...
	EditResp struct {
		// add response params
	}
)

// PostEdit ...
func (l *User) PostEdit(ctx context.Context, req *EditReq) (*EditResp, error) {
	// add your code here
	return &EditResp{}, nil
}
