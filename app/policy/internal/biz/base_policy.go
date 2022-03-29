package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
)

func (p *BasePolicy) Render() error {
	var err error
	if p.Manifest != "" {
		return nil
	}
	p.Manifest = fmt.Sprintf("base policy(%v) render yaml", p.Name)
	return err
}

// PolicyUsecase is a Greeter usecase.
type PolicyUsecase struct {
	repo    BasePolicyRepo
	temRepo ImportTemplateRepo
	log     *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewPolicyUsecase(repo BasePolicyRepo, logger log.Logger) *PolicyUsecase {
	return &PolicyUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (pc *PolicyUsecase) Create(ctx context.Context, policy *BasePolicy) (*BasePolicy, error) {
	pc.log.WithContext(ctx).Infof("CreatePolicy: %v", policy.Name)
	policy.Render()
	return pc.repo.Create(ctx, policy)
}

func (pc *PolicyUsecase) FindByID(ctx context.Context, in int64) (*BasePolicy, error) {
	_, err := pc.repo.FindByID(ctx, in)
	if err != nil {
		return nil, errors.New(
			http.StatusMethodNotAllowed,
			"Policy_NOT_FOUND",
			"未找到该用户信息",
		)
	}
	return nil, nil
}

func (pc *PolicyUsecase) Update(ctx context.Context, policy *BasePolicy) (*BasePolicy, error) {
	if err := policy.Render(); err != nil {
		return nil, err
	}
	return pc.repo.Update(ctx, policy)
}

func (pc *PolicyUsecase) ListAll(ctx context.Context, query Query) ([]*BasePolicy, error) {
	return pc.repo.ListAll(ctx, query)
}

func (pc *PolicyUsecase) Delete(ctx context.Context, in int64) (*BasePolicy, error) {
	return pc.repo.Delete(ctx, in)
}
