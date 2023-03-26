package logic

import (
	"context"
	"errors"
	"examples/mall/user/rpc/internal/svc"
	"examples/mall/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	if in.Id != "1" {
		return nil, errors.New("error without this user")
	}

	return &user.UserResponse{
		Name:   "Huang Xi",
		Gender: "男",
	}, nil
}
