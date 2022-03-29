package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-project/app/policy/internal/biz"
)

type basePolicyRepo struct {
	data *Data
	log  *log.Helper
}

// NewBasePolicyRepo .
func NewBasePolicyRepo(data *Data, logger log.Logger) biz.BasePolicyRepo {
	return &basePolicyRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *basePolicyRepo) Create(ctx context.Context, g *biz.BasePolicy) (*biz.BasePolicy, error) {
	// do someting in db client
	// r.data.db.Save()
	return g, nil
}

func (r *basePolicyRepo) Update(ctx context.Context, g *biz.BasePolicy) (*biz.BasePolicy, error) {
	return g, nil
}

func (r *basePolicyRepo) FindByID(context.Context, int64) (*biz.BasePolicy, error) {
	return nil, nil
}

func (r *basePolicyRepo) ListAll(context.Context, biz.Query) ([]*biz.BasePolicy, error) {
	return nil, nil
}

func (r *basePolicyRepo) Delete(context.Context, int64) (*biz.BasePolicy, error) {
	return nil, nil
}
