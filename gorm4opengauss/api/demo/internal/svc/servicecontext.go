package svc

import (
	"github.com/okyer/examples/gorm4opengauss/api/demo/internal/config"
	gaussdb "github.com/okyer/gorm4gaussdb"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(gaussdb.Open(c.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
