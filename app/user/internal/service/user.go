package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"kratos-project/api/user/v1"
	"kratos-project/app/user/internal/biz"
	"net/http"
)

// UserService is a greeter service.
type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUsecase
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (u *UserService) CreateUser(context.Context, *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	// do something
	// api to biz struct
	// u.uc.CreateUser()
	return nil, nil
}
func (u *UserService) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.GetUserReply, error) {
	// 可以在 biz 层直接返回 v1 的结构内容
	// 但需要该业务在 biz 层是否会发生变化，即 biz 层使用的内容与 v1 类型不完全对应
	// 或者后续的 v1 结构 是多个 biz 的复合组成
	return u.uc.GetUser(ctx, biz.User{})
}
func (u *UserService) DeleteUser(context.Context, *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	// biz 数据模型 转成 api 数据模型的逻辑，考虑到后续的组装、复用，应该写在 service 层进行对 api 接口的转换
	// 返回错误类型
	err := errors.New(http.StatusInternalServerError, "USER_NOT_FOUND", "删除失败")
	err = err.WithMetadata(map[string]string{
		"key": "layout",
	})
	return nil, err
}
func (u *UserService) ListUser(context.Context, *v1.ListUserRequest) (*v1.ListUserReply, error) {
	return nil, nil
}
