package logic

import (
	"book/common/errorx"
	"book/service/user/model"
	"context"
	"strings"

	"book/service/user/rpc/internal/svc"
	"book/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegistryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegistryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegistryLogic {
	return &RegistryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegistryLogic) Registry(req *user.RegistryReq) (*user.RegistryResp, error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 ||
		len(strings.TrimSpace(req.Gender)) == 0 || len(strings.TrimSpace(req.Number)) == 0 {
		return nil, errorx.NewDefaultError("存在请求参数为空")
	}
	u := &model.User{
		Number:   req.Number,
		Gender:   req.Gender,
		Username: req.Username,
		Password: req.Password,
	}
	if _, err := l.svcCtx.UserModel.Insert(u); err != nil {
		return &user.RegistryResp{
			Username: req.Username,
			Detail:   "注册失败",
		}, err
	}
	return &user.RegistryResp{
		Username: req.Username,
		Detail:   "注册成功",
	}, nil
}
