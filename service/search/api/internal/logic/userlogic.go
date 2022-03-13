package logic

import (
	"book/common/errorx"
	"book/service/user/rpc/user"
	"context"

	"book/service/search/api/internal/svc"
	"book/service/search/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLogic {
	return UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req types.RegistryReq) (resp *types.RegistryResp, err error) {
	// todo: add your logic here and delete this line
	request := &user.RegistryReq{
		Username: req.Username,
		Number:   req.Number,
		Password: req.Password,
		Gender:   req.Gender,
	}
	userInfo, err := l.svcCtx.UserRpc.Registry(l.ctx, request)
	if err != nil {
		return nil, errorx.NewDefaultError("rpc 注册失败")
	}
	return &types.RegistryResp{
		Detail:   "注册成功",
		Username: userInfo.Username,
	}, nil
}
