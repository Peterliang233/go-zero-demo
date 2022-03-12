package logic

import (
	"book/common/errorx"
	"context"

	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *GetUserLogic) GetUser(req types.IdReq) (*types.UserInfoReply, error) {
	userInfo, err := l.svcCtx.UserModel.FindOne(req.Id)
	if err != nil {
		return nil, errorx.NewDefaultError("参数请求错误")
	}

	return &types.UserInfoReply{
		Id:       userInfo.Id,
		Username: userInfo.Username,
		Number:   userInfo.Number,
		Gender:   userInfo.Gender,
	}, nil
}
