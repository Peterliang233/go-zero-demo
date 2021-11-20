package logic

import (
	"book/service/search/api/internal/svc"
	"book/service/search/api/internal/types"
	"book/service/user/rpc/userclient"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchLogic {
	return SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 使用rpc调用
func (l *SearchLogic) Search(req types.SearchReq) (*types.SearchReply, error) {
	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("userId: %s", userIdNumber)
	userId, err := userIdNumber.Int64()
	if err != nil {
		return nil, err
	}

	userInfo, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdReq{
		Id: userId,
	})

	if err != nil {
		return nil, err
	}

	return &types.SearchReply{
		Name:  userInfo.Name,
		Count: 100,
	}, nil
}
