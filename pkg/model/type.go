package model

type UserStatus int8

const (
	// UserStatusNormal 正常
	UserStatusNormal UserStatus = iota + 1
	// UserStatusDisabled 禁用
	UserStatusDisabled
	// UserStatusDeleted 删除
	UserStatusDeleted
)

func (s UserStatus) String() string {
	switch s {
	case UserStatusNormal:
		return "正常"
	case UserStatusDisabled:
		return "禁用"
	case UserStatusDeleted:
		return "删除"
	}
	return "未知"
}
