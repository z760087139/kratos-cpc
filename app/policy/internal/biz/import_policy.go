package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
)

func (p *ImportPolicy) Render(template string) error {
	var err error
	if p.Manifest != "" {
		return nil
	}
	p.Manifest = fmt.Sprintf("import policy(%v) render yaml", p.Name)
	return err
}

// PolicyUsecase is a Greeter usecase.
type ImportPolicyUsecase struct {
	repo    ImportPolicyRepo
	temRepo ImportTemplateRepo
	log     *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewImportPolicyUsecase(repo ImportPolicyRepo, logger log.Logger) *ImportPolicyUsecase {
	return &ImportPolicyUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (pc *ImportPolicyUsecase) Create(ctx context.Context, policy *ImportPolicy) (*ImportPolicy, error) {
	pc.log.WithContext(ctx).Infof("CreatePolicy: %v", policy.Name)
	template, err := pc.temRepo.FindById(ctx, 0)
	if err != nil {
		return nil, err
	}
	if err := policy.Render(template.BaseYaml); err != nil {
		return nil, err
	}
	return pc.repo.Create(ctx, policy)
}

func (pc *ImportPolicyUsecase) FindByID(ctx context.Context, in int64) (*ImportPolicy, error) {
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

func (pc *ImportPolicyUsecase) Update(ctx context.Context, policy *ImportPolicy) (*ImportPolicy, error) {
	template, err := pc.temRepo.FindById(ctx, 0)
	if err != nil {
		return nil, err
	}
	if err := policy.Render(template.BaseYaml); err != nil {
		return nil, err
	}
	return pc.repo.Update(ctx, policy)
}

func (pc *ImportPolicyUsecase) ListAll(ctx context.Context, query Query) ([]*ImportPolicy, error) {
	return pc.repo.ListAll(ctx, query)
}

func (pc *ImportPolicyUsecase) Delete(ctx context.Context, in int64) (*ImportPolicy, error) {
	return pc.repo.Delete(ctx, in)
}
