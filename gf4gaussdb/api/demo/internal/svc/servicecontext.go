package svc

import (
	_ "github.com/gogf/gf/contrib/drivers/gaussdb/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/okyer/examples/gf4gaussdb/api/demo/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     gdb.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	configNode := gdb.ConfigNode{
		Type: "gaussdb",
		// Host:  "127.0.0.1",
		// Port:  "8000",
		// User:  "test",
		// Pass:  "123456",
		// Name:  "test",
		// Extra: "sslmode=disable",
		Link: c.DSN,
	}
	gdb.AddConfigNode(gdb.DefaultGroupName, configNode)
	db, err := gdb.New(configNode)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
