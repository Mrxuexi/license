package svc

import (
	"examples/greet/internal/config"
)

// ServiceContext 服务上下文
type ServiceContext struct {
	Config config.Config // 嵌套了 restful api config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
