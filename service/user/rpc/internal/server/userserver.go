// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"book/service/user/rpc/internal/logic"
	"book/service/user/rpc/internal/svc"
	"book/service/user/rpc/user"
	"context"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user.IdReq) (*user.UserInfoReply, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginReq) (*user.LoginReply, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Registry(ctx context.Context, in *user.RegistryReq) (*user.RegistryResp, error) {
	l := logic.NewRegistryLogic(ctx, s.svcCtx)
	return l.Registry(in)
}
