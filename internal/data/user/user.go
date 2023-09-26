package user

import (
	"gin-plus-admin/pkg/conn"
	"gin-plus-admin/pkg/model"
	"gin-plus-admin/pkg/model/query"

	"gorm.io/gorm"
)

type (
	// User ...
	User struct {
		query.IAction[model.User]

		PreloadFilesKey string
	}
)

// NewUser ...
func NewUser() *User {
	return &User{
		IAction:         query.NewAction(query.WithDB[model.User](conn.GetMysqlDB())),
		PreloadFilesKey: model.PreloadFiles,
	}
}

// PreloadFiles 预加载关联文件
func (l *User) PreloadFiles(scops ...model.Scopemethod) model.Scopemethod {
	return func(db *gorm.DB) *gorm.DB {
		if len(scops) == 0 {
			return db.Preload(l.PreloadFilesKey)
		}
		// add your code here
		return db.Preload(l.PreloadFilesKey, func(db *gorm.DB) *gorm.DB {
			return db.Scopes(scops...)
		})
	}
}
