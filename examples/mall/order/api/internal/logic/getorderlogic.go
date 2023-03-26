package logic

import (
	"context"
	"examples/mall/user/rpc/types/user"

	"examples/mall/order/api/internal/svc"
	"examples/mall/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {

	userResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.OrderReply{
		Id:     userResp.GetId(),
		Name:   userResp.GetName(),
		Gender: userResp.GetGender(),
	}, nil
}
