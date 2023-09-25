package role

import (
	"context"
)

type (
	// UpdateStatusReq ...
	UpdateStatusReq struct {
		// add request params
		ID int `uri:"id"`
	}

	// UpdateStatusResp ...
	UpdateStatusResp struct {
		// add response params
	}
)

// UpdateStatus ...
func (l *Role) UpdateStatus(ctx context.Context, req *UpdateStatusReq) (*UpdateStatusResp, error) {
	// add your code here
	return &UpdateStatusResp{}, nil
}
