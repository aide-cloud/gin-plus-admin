package user

import (
	"context"
	dataUser "gin-plus-admin/internal/data/user"

	ginplus "github.com/aide-cloud/gin-plus"
	query "github.com/aide-cloud/gorm-normalize"
)

type (
	// DeleteReq ...
	DeleteReq struct {
		// add request params
		ID uint `uri:"id"`
	}

	// DeleteResp ...
	DeleteResp struct {
		// add response params
	}
)

// Delete ...
func (l *User) Delete(ctx context.Context, req *DeleteReq) (*DeleteResp, error) {
	userData := dataUser.NewUser()

	if err := userData.Delete(query.WhereID(req.ID)); err != nil {
		ginplus.Logger().Sugar().Error(err)
		return nil, err
	}
	// add your code here
	return &DeleteResp{}, nil
}
