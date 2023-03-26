// main 一个简单的 restful api 服务
// curl -i -X GET http://localhost:8888/from/you 请求一下示例路由返回 Hello go-zero
package main

import (
	"flag"
	"fmt"

	"examples/greet/internal/config"
	"examples/greet/internal/handler"
	"examples/greet/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/greet-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
