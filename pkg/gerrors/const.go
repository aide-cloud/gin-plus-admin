package gerrors

import (
	"errors"
)

type Code int

const (
	// Unknown 未知错误
	Unknown Code = iota + 1000
	// ErrSystem 系统错误
	ErrSystem
	// ErrInternal 内部错误
	ErrInternal
	// ErrInvalidParam 参数错误
	ErrInvalidParam
	// ErrInvalidToken 无效的token
	ErrInvalidToken
	// ErrInvalidUser 无效的用户
	ErrInvalidUser
)

func (c Code) Error() error {
	switch c {
	case ErrSystem:
		return errors.New("系统错误")
	case ErrInternal:
		return errors.New("内部错误")
	case ErrInvalidParam:
		return errors.New("参数错误")
	case ErrInvalidToken:
		return errors.New("无效的token")
	case ErrInvalidUser:
		return errors.New("无效的用户")
	default:
		return errors.New("未知错误")
	}
}
