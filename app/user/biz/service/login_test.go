package service

import (
	"context"
	"gomall/app/user/biz/dal/mysql"
	user "gomall/rpc_gen/kitex_gen/user"
	"testing"

	"github.com/joho/godotenv"
)

func TestLogin_Run(t *testing.T) {
	// 加载环境变量
	godotenv.Load("../../.env")

	// 初始化数据库连接
	mysql.Init()

	ctx := context.Background()
	s := NewLoginService(ctx)

	// 测试用例
	tests := []struct {
		name    string
		req     *user.LoginReq
		wantErr bool
	}{
		{
			name: "valid login",
			req: &user.LoginReq{
				Email:    "test@example.com",
				Password: "123456",
			},
			wantErr: false,
		},
		{
			name: "empty email",
			req: &user.LoginReq{
				Email:    "",
				Password: "123456",
			},
			wantErr: true,
		},
		{
			name: "empty password",
			req: &user.LoginReq{
				Email:    "test@example.com",
				Password: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := s.Run(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && resp != nil {
				if resp.UserId == 0 {
					t.Error("LoginService.Run() returned zero user ID for valid login")
				}
				if resp.Token == "" {
					t.Error("LoginService.Run() returned empty token for valid login")
				}
			}
		})
	}
}
