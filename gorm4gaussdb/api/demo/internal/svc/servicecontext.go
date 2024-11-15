package svc

import (
	_ "gitee.com/opengauss/openGauss-connector-go-pq"
	"github.com/okyer/examples/gorm4gaussdb/api/demo/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(postgres.New(postgres.Config{DriverName: "postgres", DSN: c.DSN}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
