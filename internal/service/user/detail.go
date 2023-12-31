package user

import (
	"context"

	dataUser "gin-plus-admin/internal/data/user"
	"gin-plus-admin/pkg/model"

	ginplus "github.com/aide-cloud/gin-plus"
	query "github.com/aide-cloud/gorm-normalize"
	"go.uber.org/zap"
)

type (
	// DetailReq ...
	DetailReq struct {
		// add request params
		ID uint `uri:"id"`
	}

	// DetailResp ...
	DetailResp struct {
		// add response params
		ID        uint   `json:"id"`
		Name      string `json:"name"`
		CreatedAt int64  `json:"created_at"`
		UpdateAt  int64  `json:"update_at"`

		Files []*model.File `json:"files"`
	}
)

// GetDetail ...
func (l *User) GetDetail(ctx context.Context, req *DetailReq) (*DetailResp, error) {
	userData := dataUser.NewUser()

	first, err := userData.WithContext(ctx).First(query.WhereID(req.ID), userData.PreloadFiles())
	if err != nil {
		ginplus.Logger().Error("get user detail failed", zap.Any("req", req), zap.Error(err))
		return nil, err
	}

	// add your code here
	return &DetailResp{
		ID:        first.ID,
		Name:      first.Name,
		CreatedAt: first.CreatedAt.Unix(),
		UpdateAt:  first.UpdatedAt.Unix(),
		Files:     first.Files,
	}, nil
}
