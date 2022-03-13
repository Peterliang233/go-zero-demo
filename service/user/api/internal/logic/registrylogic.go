package logic

import (
	"book/common/errorx"
	"book/service/user/model"
	"context"
	"net/http"
	"strings"

	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"

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
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 ||
		len(strings.TrimSpace(req.Gender)) == 0 || len(strings.TrimSpace(req.Number)) == 0 {
		return nil, errorx.NewDefaultError("存在请求参数为空")
	}
	user := model.User{
		Username: req.Username,
		Number:   req.Number,
		Gender:   req.Gender,
		Password: req.Password,
	}
	if _, err := l.svcCtx.UserModel.Insert(user); err != nil {
		return nil, errorx.NewCodeError(http.StatusInternalServerError, "数据库执行错误")
	}
	return &types.RegistryResp{
		Username: req.Username,
		Number:   req.Number,
		Gender:   req.Gender,
		Detail:   "注册成功",
	}, nil
}
