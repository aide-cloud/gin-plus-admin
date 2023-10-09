package model

import (
	query "github.com/aide-cloud/gorm-normalize"
	"gorm.io/gorm"
)

type User struct {
	query.BaseModel
	Name     string     `gorm:"column:name;type:varchar(20);not null;default:'';comment:用户名" json:"name"`
	Password string     `gorm:"column:password;type:varchar(20);not null;default:'';comment:用户密码" json:"-"`
	Salt     string     `gorm:"column:salt;type:char(8);not null;default:'';comment:加密盐值" json:"-"`
	Birthday uint       `gorm:"column:birthday;type:int unsigned;default:0;NOT NULL;comment:生日20060102" json:"birthday"`
	Avatar   string     `gorm:"column:avatar;type:varchar(255);default:'';NOT NULL;comment:用户当前头像" json:"avatar"`
	Status   UserStatus `gorm:"column:status;type:tinyint unsigned;default:0;NOT NULL;comment:用户状态" json:"status"`

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

// 查询之后
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	if u == nil {
		return
	}

	if u.DeletedAt > 0 {
		u.Status = UserStatusDeleted
	}

	return
}
