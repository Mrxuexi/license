package logic

// 导入需要使用的包
import (
	"context"
	"examples/greet/internal/svc"            // 服务包，用于访问其他服务
	"examples/greet/internal/types"          // 定义请求和响应的数据类型包
	"github.com/zeromicro/go-zero/core/logx" // 日志包，用于记录日志
)

// GreetLogic 业务结构体，包含日志、上下文和服务上下文
type GreetLogic struct {
	logx.Logger                     // 继承 Logger，用于日志记录
	ctx         context.Context     // 上下文
	svcCtx      *svc.ServiceContext // 服务上下文
}

// NewGreetLogic 创建 GreetLogic 实例，参数包括上下文和服务上下文
func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	// 返回 GreetLogic 实例
	return &GreetLogic{
		Logger: logx.WithContext(ctx), // 初始化日志记录器
		ctx:    ctx,                   // 初始化上下文
		svcCtx: svcCtx,                // 初始化服务上下文
	}
}

// Greet 方法，接收请求数据类型的指针并返回响应数据类型和错误
func (l *GreetLogic) Greet(req *types.Request) (resp *types.Response, err error) {
	// 返回响应数据类型
	return &types.Response{
		Message: "Hello go-zero",
	}, nil
}
