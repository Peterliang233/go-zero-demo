package logic

import (
	"book/common/errorx"
	"book/service/user/model"
	"book/service/user/rpc/internal/svc"
	"book/service/user/rpc/user"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginReply, error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByNumber(in.Username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("用户不存在")
	default:
		return nil, err
	}

	if userInfo.Password != in.Password {
		return nil, errorx.NewDefaultError("用户密码错误")
	}

	return &user.LoginReply{
		Id:     userInfo.Id,
		Name:   userInfo.Name,
		Gender: userInfo.Gender,
	}, nil
}
