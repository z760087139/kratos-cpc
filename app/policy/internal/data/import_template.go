package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-project/api/importTemplate/v1"
	"kratos-project/app/policy/internal/biz"
	"kratos-project/app/policy/internal/data/internal"
)

type importPolicyRepo struct {
	data      *Data
	tableName string
	log       *log.Helper
}

// NewimportPolicyRepo .
func NewImportPolicyRepo(data *Data, logger log.Logger) biz.ImportPolicyRepo {
	return &importPolicyRepo{
		data:      data,
		tableName: "import_policy",
		log:       log.NewHelper(logger),
	}
}

func (r *importPolicyRepo) Create(ctx context.Context, g *biz.ImportPolicy) (*biz.ImportPolicy, error) {
	// do someting in db client
	// r.data.db.Save()
	return g, nil
}

func (r *importPolicyRepo) Update(ctx context.Context, g *biz.ImportPolicy) (*biz.ImportPolicy, error) {
	return g, nil
}

func (r *importPolicyRepo) FindByID(ctx context.Context, in int64) (*biz.ImportPolicy, error) {
	var data internal.ImportPolicy
	r.data.db.Table(r.tableName).First(data, in)
	return data.FromInternal()
}

func (r *importPolicyRepo) ListAll(context.Context, biz.Query) ([]*biz.ImportPolicy, error) {
	return nil, nil
}

func (r *importPolicyRepo) Delete(context.Context, int64) (*biz.ImportPolicy, error) {
	return nil, nil
}

type importTemplateRepo struct {
	data *Data
	log  *log.Helper
}

func NewImportTemplateRepo(data *Data, logger log.Logger) biz.ImportTemplateRepo {
	return &importTemplateRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *importTemplateRepo) FindById(ctx context.Context, in int64) (biz.Template, error) {
	template, err := r.data.tmCli.GetPolicy(ctx, &v1.GetTemplateRequest{Id: in})
	if err != nil {
		return biz.Template{}, err
	}
	return biz.Template{
		BaseYaml: template.GetYaml(),
	}, nil
}
