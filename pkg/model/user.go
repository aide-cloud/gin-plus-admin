package model

type User struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(20);not null;default:'';comment:用户名" json:"name"`
}

const (
	tableName = "users"
)

func (User) TableName() string {
	return tableName
}
