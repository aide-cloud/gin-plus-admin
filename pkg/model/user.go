package model

import query "github.com/aide-cloud/gorm-normalize"

type User struct {
	query.BaseModel
	Name  string  `gorm:"column:name;type:varchar(20);not null;default:'';comment:用户名" json:"name"`
	Files []*File `gorm:"foreignKey:UserID" json:"files"`
}

const (
	userTableName = "users"
)

const (
	// PreloadFiles 预加载关联文件
	PreloadFiles = "Files"
)

func (User) TableName() string {
	return userTableName
}
