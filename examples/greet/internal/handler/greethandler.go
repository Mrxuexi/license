package handler

// 导入需要使用的包
import (
	"examples/greet/internal/logic"           // 逻辑包，处理请求和生成响应
	"examples/greet/internal/svc"             // 服务包，用于访问其他服务
	"examples/greet/internal/types"           // 定义请求和响应的数据类型包
	"github.com/zeromicro/go-zero/rest/httpx" // HTTP 工具包，用于处理 HTTP 请求和响应
	"net/http"                                // HTTP 包，用于处理 HTTP 请求和响应
)

// GreetHandler 方法，接收服务上下文并返回 http.HandlerFunc 类型的方法
func GreetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request

		// 使用 httpx.Parse 方法将请求解析到请求数据类型中，并检查解析过程中是否出错
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGreetLogic(r.Context(), svcCtx) // 实例化 GreetLogic 类型
		resp, err := l.Greet(&req)                    // 调用 Greet 方法，生成响应数据和错误
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
