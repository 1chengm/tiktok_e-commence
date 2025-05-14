package service

import (
	"context"
	"errors"
	"fmt"
	"gomall/app/user/biz/dal/mysql"
	"gomall/app/user/biz/model"
	"gomall/app/user/utils"
	user "gomall/rpc_gen/kitex_gen/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginService struct {
	ctx context.Context
}

// NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// 参数验证
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}

	// 查找用户
	row, err := model.GetByEmail(mysql.DB, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// 生成新的 JWT Token
	tokenString, err := utils.GenerateJWT(int32(row.ID))
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	// 返回成功响应
	resp = &user.LoginResp{
		UserId:  int32(row.ID),
		Message: "login success",
		Token:   tokenString,
	}
	return resp, nil
}
