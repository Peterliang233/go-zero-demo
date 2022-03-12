package logic

import (
	"book/common/errorx"
	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"
	"book/service/user/model"
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
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

// 使用rpc调用
//func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginReply, error) {
//	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
//		return nil, errorx.NewDefaultError("参数错误")
//	}
//
//	userInfo, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginReq{
//		Username: req.Username,
//		Password: req.Password,
//	})
//
//	now := time.Now().Unix()
//
//	accessExpire := l.svcCtx.Config.Auth.AccessExpire
//
//	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return &types.LoginReply{
//		Id:           userInfo.Id,
//		Username:     userInfo.Name,
//		Gender:       userInfo.Gender,
//		AccessToken:  jwtToken,
//		AccessExpire: now + accessExpire,
//		RefreshAfter: now + accessExpire/2,
//	}, nil
//}

// 使用api调用
func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginReply, error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
	}

	userInfo, err := l.svcCtx.UserModel.FindOneByUsername(req.Username)

	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("用户不存在")
	default:
		return nil, err
	}

	if userInfo.Password != req.Password {
		return nil, errors.New("用户密码不正确")
	}

	now := time.Now().Unix()

	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)

	if err != nil {
		return nil, err
	}

	return &types.LoginReply{
		Id:           userInfo.Id,
		Username:     userInfo.Username,
		Gender:       userInfo.Gender,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
