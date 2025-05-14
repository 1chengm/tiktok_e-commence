package service

import (
	"context"
	"errors"
	"gomall/app/user/biz/dal/mysql"
	"gomall/app/user/biz/model"
	user "gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gomall/app/user/utils"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" || req.PasswordConfirm == "" {
		return nil, errors.New("email or password or password confirm is empty")
	}
	if req.Password != req.PasswordConfirm {
		return nil , errors.New("password and password confirm not match")
	}
	 // 检查用户是否已存在 (推荐添加此逻辑)
    _, err = model.GetByEmail(mysql.DB, req.Email)
    if err == nil { // 如果找到了用户，说明邮箱已注册
    	return nil, errors.New("email already registered")
    } else if !errors.Is(err, gorm.ErrRecordNotFound) { // 如果是其他数据库错误
    	return nil, err
    }
	passwordHashed ,err:= bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Email: req.Email,
		PasswordHashed: string(passwordHashed),
	}
	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}
	  // 生成 JWT Token
    tokenString, err := utils.GenerateJWT(int32(newUser.ID)) // 假设 GenerateJWT 函数已定义并可用
    if err != nil {
        // 处理 Token 生成失败的情况，可以考虑返回错误或者不返回 Token 但标记注册成功
        // 为简单起见，这里直接返回错误
        return nil, errors.New("failed to generate token after registration")
    }
	return &user.RegisterResp{
		UserId: int32(newUser.ID),
		Token: tokenString,	
		Email: newUser.Email,
		}, nil
	
}
