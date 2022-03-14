package logic

import (
	"book/common/errorx"
	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"
	"book/service/user/rpc/userclient"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req types.IdReq) (resp *types.UserInfoReply, err error) {
	userInfo, err := l.svcCtx.UserRpc.GetUser(l.ctx,
		&userclient.IdReq{
			Id: req.Id,
		})
	if err != nil {
		return nil, errorx.NewDefaultError("获取用户信息失败,请检查后重试")
	}
	return &types.UserInfoReply{
		Id:       userInfo.GetId(),
		Number:   userInfo.GetNumber(),
		Username: userInfo.GetName(),
		Gender:   userInfo.GetGender(),
	}, nil
}
