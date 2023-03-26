package svc

import (
	"examples/mall/order/api/internal/config"
	"examples/mall/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User // UserRpc rpc 客户端请求参数
}

func NewServiceContext(c config.Config) *ServiceContext {
	zrpc.MustNewClient(c.UserRpc).Conn()
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
