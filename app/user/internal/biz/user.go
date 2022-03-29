package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-project/api/user/v1"
	"net/http"
)

type User struct {
	ID   int64
	Name string
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	ListByHello(context.Context, string) ([]*User, error)
	ListAll(context.Context) ([]*User, error)
}

// UserUsecase is a Greeter usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *UserUsecase) CreateUser(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

func (uc *UserUsecase) GetUser(ctx context.Context, in User) (*v1.GetUserReply, error) {
	_, err := uc.repo.FindByID(ctx, in.ID)
	if err != nil {
		return nil, errors.New(
			http.StatusMethodNotAllowed,
			"USER_NOT_FOUND",
			"未找到该用户信息",
		)
	}
	return nil, nil
}
