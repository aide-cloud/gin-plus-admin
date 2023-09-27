package user

import (
	"context"
	"errors"
	"unicode/utf8"

	dataUser "gin-plus-admin/internal/data/user"
	"gin-plus-admin/pkg/gerrors"
	"gin-plus-admin/pkg/model"
	"gin-plus-admin/pkg/util"
	"gorm.io/gorm"

	ginplus "github.com/aide-cloud/gin-plus"
)

var _ ginplus.IValidater = (*CreateReq)(nil)

type (
	// CreateReq ...
	CreateReq struct {
		// add request params
		Name     string `json:"name"`
		Password string `json:"password"`
		Birthday uint   `json:"birthday"`
		Avatar   string `json:"avatar"`
	}

	// CreateResp ...
	CreateResp struct {
		// add response params
		ID uint `json:"id"`
	}
)

var (
	ErrNameEmpty     = gerrors.New().AddFields("name", "name is empty").WithCode(gerrors.ErrInvalidParam).WithMessage("name is empty")
	ErrPasswordEmpty = gerrors.New().AddFields("password", "password is empty").WithCode(gerrors.ErrInvalidParam).WithMessage("password is empty")
	ErrorUserExist   = gerrors.New().AddFields("name", "user is exist").WithCode(gerrors.ErrInvalidParam).WithMessage("user is exist")
	ErrNameTooLong   = gerrors.New().AddFields("name", "name is too long").WithCode(gerrors.ErrInvalidParam).WithMessage("name is too long")
	ErrNameTooShort  = gerrors.New().AddFields("name", "name is too short").WithCode(gerrors.ErrInvalidParam).WithMessage("name is too short")
)

func (l *CreateReq) Validate() error {
	err := gerrors.NewNil().WithError(l.checkName()).WithError(l.checkPassword()).IsNil()
	if err != nil {
		ginplus.Logger().Sugar().Error(err)
		return err
	}

	return nil
}

func (l *CreateReq) checkName() error {
	if l.Name == "" {
		return ErrNameEmpty
	}

	if utf8.RuneCountInString(l.Name) > 20 {
		return ErrNameTooLong
	}

	if utf8.RuneCountInString(l.Name) < 2 {
		return ErrNameTooShort
	}

	// check name is exist
	userData := dataUser.NewUser()
	if _, err := userData.First(userData.WhereEqualName(l.Name)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	return ErrorUserExist
}

func (l *CreateReq) checkPassword() error {
	if l.Password == "" {
		return ErrPasswordEmpty
	}
	return nil
}

// Create ...
func (l *User) Create(ctx context.Context, req *CreateReq) (*CreateResp, error) {
	userData := dataUser.NewUser()

	// 生成8位hash盐值
	salt := util.GenerateSalt()
	// 加密密码
	password, err := util.EncryptPassword(req.Password, salt)
	if err != nil {
		ginplus.Logger().Sugar().Errorf("encryptPassword error: %v", err)
		return nil, err
	}

	newUserInfo := model.User{
		Name:     req.Name,
		Password: password,
		Birthday: req.Birthday,
		Avatar:   req.Avatar,
		Salt:     salt,
	}
	if err = userData.WithContext(ctx).Create(&newUserInfo); err != nil {
		ginplus.Logger().Sugar().Errorf("create user error: %v", err)
		return nil, err
	}
	// add your code here
	return &CreateResp{
		ID: newUserInfo.ID,
	}, nil
}
