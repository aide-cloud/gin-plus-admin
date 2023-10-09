package user

import (
	"context"
	dataUser "gin-plus-admin/internal/data/user"
	"gin-plus-admin/pkg/model"
)

type (
	// UpdateStatusReq ...
	UpdateStatusReq struct {
		// add request params
		ID     uint             `uri:"id"`
		Status model.UserStatus `json:"status"`
	}

	// UpdateStatusResp ...
	UpdateStatusResp struct {
		// add response params
	}
)

// UpdateStatus ...
func (l *User) UpdateStatus(ctx context.Context, req *UpdateStatusReq) (*UpdateStatusResp, error) {
	userData := dataUser.NewUser()

	if err := userData.Update(&model.User{Status: req.Status}); err != nil {
		return nil, err
	}

	// add your code here
	return &UpdateStatusResp{}, nil
}
