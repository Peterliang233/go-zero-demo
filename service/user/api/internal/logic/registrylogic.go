package logic

import (
	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"
	"context"

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
	// todo: add your logic here and delete this line

	return
}
