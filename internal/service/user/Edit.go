package user

import (
	"context"
	dataUser "gin-plus-admin/internal/data/user"
	"gin-plus-admin/pkg/model"

	ginplus "github.com/aide-cloud/gin-plus"
	query "github.com/aide-cloud/gorm-normalize"
)

type (
	// EditReq ...
	EditReq struct {
		// add request params
		ID       uint   `uri:"id"`
		Name     string `json:"name"`
		Birthday uint   `json:"birthday"`
		Avatar   string `json:"avatar"`
	}

	// EditResp ...
	EditResp struct {
		// add response params
	}
)

// PostEdit ...
func (l *User) PostEdit(ctx context.Context, req *EditReq) (*EditResp, error) {
	userData := dataUser.NewUser()

	if err := userData.Update(&model.User{
		Name:     req.Name,
		Birthday: req.Birthday,
		Avatar:   req.Avatar,
	}, query.WhereID(req.ID)); err != nil {
		ginplus.Logger().Sugar().Error(err)
		return nil, err
	}
	// add your code here
	return &EditResp{}, nil
}
