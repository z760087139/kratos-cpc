package data

import (
	"context"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"gorm.io/gorm"
	v1 "kratos-project/api/importTemplate/v1"
	"kratos-project/app/policy/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDiscovery,
	NewDBClient, NewImportTemplateClient,
	NewImportPolicyRepo, NewImportTemplateRepo, NewBasePolicyRepo)

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

// Data .
type Data struct {
	db    *gorm.DB
	tmCli v1.ImportTemplateClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, tmCli v1.ImportTemplateClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:    db,
		tmCli: tmCli,
	}, cleanup, nil
}

func NewDBClient(conf *conf.Data, logger log.Logger) (*gorm.DB, error) {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//return gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return &gorm.DB{}, nil
}

func NewImportTemplateClient(r registry.Discovery) v1.ImportTemplateClient {
	// reTest()
	return v1.NewImportTemplateClient(nil)
}

func reTest(r registry.Discovery) v1.ImportTemplateClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///beer.catalog.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			//tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewImportTemplateClient(conn)
}
