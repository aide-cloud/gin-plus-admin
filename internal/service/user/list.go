package user

import (
	"context"
	dataUser "gin-plus-admin/internal/data/user"
	"gin-plus-admin/pkg/conn"
	"gin-plus-admin/pkg/model"

	ginplus "github.com/aide-cloud/gin-plus"
	query "github.com/aide-cloud/gorm-normalize"
)

type (
	// ListReq ...
	ListReq struct {
		// add request params
		Keyword string `form:"keyword"`
		Curr    int    `form:"curr"`
		Size    int    `form:"size"`
	}

	// ListResp ...
	ListResp struct {
		// add response params
		List     []*model.User `json:"list"`
		PageInfo *query.Page   `json:"page_info"`
	}
)

// PostList ...
func (l *User) List(ctx context.Context, req *ListReq) (*ListResp, error) {
	userData := dataUser.NewUser()

	pgInfo := query.NewPage(req.Curr, req.Size)

	// 查询全部用户, 包括软删除的
	list, err := userData.WithDB(conn.GetMysqlDB().Unscoped()).List(pgInfo, query.WhereLikeKeyword(req.Keyword, "name"))
	if err != nil {
		ginplus.Logger().Sugar().Error(err)
		return nil, err
	}
	// add your code here
	return &ListResp{
		List:     list,
		PageInfo: pgInfo,
	}, nil
}
