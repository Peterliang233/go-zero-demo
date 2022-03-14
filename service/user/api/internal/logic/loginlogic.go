package logic

import (
	"book/common/errorx"
	"book/service/user/rpc/userclient"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"

	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginReply, err error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.NewCodeError(400, "参数不能为空")
	}
	request := &userclient.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}
	UserInfo, err := l.svcCtx.UserRpc.Login(l.ctx, request)
	if err != nil {
		return nil, errorx.NewCodeError(500, "登录错误")
	}
	now := time.Now().Unix()
	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret,
		now, l.svcCtx.Config.Auth.AccessExpire, UserInfo.GetId())
	if err != nil {
		return nil, errorx.NewCodeError(500, "获取token失败")
	}
	return &types.LoginReply{
		Username:     UserInfo.Name,
		Gender:       UserInfo.Gender,
		AccessToken:  token,
		AccessExpire: now + l.svcCtx.Config.Auth.AccessExpire,
		RefreshAfter: now + l.svcCtx.Config.Auth.AccessExpire/2,
	}, nil
}

func (l *LoginLogic) getJwtToken(secret string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}
