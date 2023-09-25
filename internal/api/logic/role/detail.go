package role

import (
	"context"
)

type (
	// DetailReq ...
	DetailReq struct {
		// add request params
		ID int `uri:"id"`
	}

	// DetailResp ...
	DetailResp struct {
		// add response params
	}
)

// Detail ...
func (l *Role) Detail(ctx context.Context, req *DetailReq) (*DetailResp, error) {
	// add your code here
	return &DetailResp{}, nil
}
