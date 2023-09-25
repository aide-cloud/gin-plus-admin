package user

import (
	"context"

	"gin-plus-admin/pkg/conn"
	"gin-plus-admin/pkg/model"
	"gin-plus-admin/pkg/model/methods"

	ginplus "github.com/aide-cloud/gin-plus"
	"go.uber.org/zap"
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
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
)

// GetDetail ...
func (l *User) GetDetail(ctx context.Context, req *DetailReq) (*DetailResp, error) {
	action := methods.NewAction(methods.WithDB[model.User](conn.GetMysqlDB()))

	first, err := action.First(ctx, model.WhereID(req.ID))
	if err != nil {
		ginplus.Logger().Error("get user detail failed", zap.Any("req", req), zap.Error(err))
		return nil, err
	}

	// add your code here
	return &DetailResp{
		ID:   first.ID,
		Name: first.Name,
	}, nil
}
