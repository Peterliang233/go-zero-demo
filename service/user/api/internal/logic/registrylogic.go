package logic

import (
	"book/common/errorx"
	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"
	"book/service/user/rpc/userclient"
	"context"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegistryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegistryLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegistryLogic {
	return RegistryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegistryLogic) Registry(req types.RegistryReq) (resp *types.RegistryResp, err error) {
	if len(strings.TrimSpace(req.Password)) == 0 || len(strings.TrimSpace(req.Username)) == 0 ||
		len(strings.TrimSpace(req.Gender)) == 0 || len(strings.TrimSpace(req.Number)) == 0 {
		return nil, errorx.NewDefaultError("请求参数不能为空")
	}

	user := &userclient.RegistryReq{
		Username: req.Username,
		Password: req.Password,
		Number:   req.Number,
		Gender:   req.Gender,
	}
	userInfo, err := l.svcCtx.UserRpc.Registry(l.ctx, user)
	if err != nil {
		return nil, errorx.NewDefaultError("api注册失败")
	}
	return &types.RegistryResp{
		Username: userInfo.GetUsername(),
		Detail:   userInfo.GetDetail(),
		Gender:   req.Gender,
		Number:   req.Number,
	}, nil
}
