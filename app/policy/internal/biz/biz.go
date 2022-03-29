package biz

import (
	"context"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewPolicyUsecase, NewImportPolicyUsecase)

// BasePolicyRepo is a Greater repo.
type BasePolicyRepo interface {
	Create(context.Context, *BasePolicy) (*BasePolicy, error)
	Update(context.Context, *BasePolicy) (*BasePolicy, error)
	FindByID(context.Context, int64) (*BasePolicy, error)
	ListAll(context.Context, Query) ([]*BasePolicy, error)
	Delete(context.Context, int64) (*BasePolicy, error)
}

type ImportTemplateRepo interface {
	FindById(context.Context, int64) (Template, error)
}

// BasePolicyRepo is a Greater repo.
type ImportPolicyRepo interface {
	Create(context.Context, *ImportPolicy) (*ImportPolicy, error)
	Update(context.Context, *ImportPolicy) (*ImportPolicy, error)
	FindByID(context.Context, int64) (*ImportPolicy, error)
	ListAll(context.Context, Query) ([]*ImportPolicy, error)
	Delete(context.Context, int64) (*ImportPolicy, error)
}
