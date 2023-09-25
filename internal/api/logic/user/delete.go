package user

import (
	"context"
)

type (
	// DeleteReq ...
	DeleteReq struct {
		// add request params
		ID int `uri:"id"`
	}

	// DeleteResp ...
	DeleteResp struct {
		// add response params
	}
)

// Delete ...
func (l *User) Delete(ctx context.Context, req *DeleteReq) (*DeleteResp, error) {
	// add your code here
	return &DeleteResp{}, nil
}
